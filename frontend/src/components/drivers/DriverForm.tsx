import React, { useEffect, useState } from 'react';
import { Form, Input, DatePicker, Select, Button, Space, message, Image, Row, Col } from 'antd';
import { PlusOutlined, DeleteOutlined } from '@ant-design/icons';
import dayjs from 'dayjs';
import {  useCreateDriver,
  useUpdateDriver,
  useUploadDriverPhoto,
  useDeleteDriverPhoto,
  useUploadDriverLicense,
  useUploadDriverPassport,
  useDeleteDriverLicense,
  useDeleteDriverPassport,
} from '../../hooks/useDrivers';
import type { Driver, CreateDriverRequest, UpdateDriverRequest } from '../../types/driver';
const { Option } = Select;
interface DriverFormProps {
  driver?: Driver;
  onSuccess: () => void;
  onCancel: () => void;
}

export const DriverForm: React.FC<DriverFormProps> = ({ driver, onSuccess, onCancel }) => {
  const [form] = Form.useForm();
  const createDriver = useCreateDriver();
  const updateDriver = useUpdateDriver();
  const uploadPhoto = useUploadDriverPhoto();
  const deletePhoto = useDeleteDriverPhoto();
  const uploadLicense = useUploadDriverLicense();
  const uploadPassport = useUploadDriverPassport();
  const deleteLicense = useDeleteDriverLicense();
  const deletePassport = useDeleteDriverPassport();
  const [photoUrl, setPhotoUrl] = useState<string | undefined>(driver?.photo);
  const [licenseScanUrl, setLicenseScanUrl] = useState<string | undefined>(driver?.driverLicenseScan);
  const [passportScanUrl, setPassportScanUrl] = useState<string | undefined>(driver?.passportScan);

  const [uploading, setUploading] = useState(false);
  const [tempPhotoFile, setTempPhotoFile] = useState<File | null>(null);
  const isEdit = !!driver;

  useEffect(() => {
    if (driver) {
      form.setFieldsValue({
        ...driver,
        birthDate: driver.birthDate ? dayjs(driver.birthDate) : null,
        driverLicenseIssueDate: dayjs(driver.driverLicenseIssueDate),
        driverLicenseExpiryDate: dayjs(driver.driverLicenseExpiryDate),
        passportIssueDate: driver.passportIssueDate ? dayjs(driver.passportIssueDate) : null,
      });
      setPhotoUrl(driver.photo);
      setLicenseScanUrl(driver.driverLicenseScan);
      setPassportScanUrl(driver.passportScan);
    }
  }, [driver, form]);

const handlePhotoDelete = async () => {
    if (isEdit && driver && driver.photo) {
      try {
        await deletePhoto.mutateAsync(driver.id);
        message.success('Фото успешно удалено');
      } catch (error) {
        message.error('Ошибка при удалении фото');
        return;
      }
    }
    setPhotoUrl(undefined);
    setTempPhotoFile(null);
    form.setFieldValue('photo', undefined);
  };
  const handleLicenseDelete = async () => {
    if (isEdit && driver && driver.driverLicenseScan) {
      try {
        await deleteLicense.mutateAsync(driver.id);
        message.success('Скан прав успешно удален');
      } catch (error) {
        message.error('Ошибка при удалении скана прав');
        return;
      }
    }
    setLicenseScanUrl(undefined);
    form.setFieldValue('driverLicenseScan', undefined);
  };

  const handlePassportDelete = async () => {
    if (isEdit && driver && driver.passportScan) {
      try {
        await deletePassport.mutateAsync(driver.id);
        message.success('Скан паспорта успешно удален');
      } catch (error) {
        message.error('Ошибка при удалении скана паспорта');
        return;
      }
    }
    setPassportScanUrl(undefined);
    form.setFieldValue('passportScan', undefined);
  };
  const beforeUpload = (file: File) => {
    const isImage = file.type.startsWith('image/');
    if (!isImage) {
      message.error('Можно загружать только изображения!');
      return false;
    }
    const isLt5M = file.size / 1024 / 1024 < 5;
    if (!isLt5M) {
      message.error('Изображение должно быть меньше 5MB!');
      return false;
    }
    return true;
  };

  const customRequest = async (options: any) => {
    const { file, onSuccess, onError, type } = options;

    if (isEdit && driver) {
      try {
        setUploading(true);
        let result;

        if (type === 'licenseScan') {
          result = await uploadLicense.mutateAsync({ id: driver.id, file, fileType: 'scan' });
          setLicenseScanUrl(result.filePath);
          form.setFieldValue('driverLicenseScan', result.filePath);
        } else if (type === 'passportScan') {
          result = await uploadPassport.mutateAsync({ id: driver.id, file, fileType: 'scan' });
          setPassportScanUrl(result.filePath);
          form.setFieldValue('passportScan', result.filePath);
        } else {
          result = await uploadPhoto.mutateAsync({ id: driver.id, file });
          setPhotoUrl(result.filePath);
          form.setFieldValue('photo', result.filePath);
        }

        onSuccess(result);
        message.success('Файл успешно загружен');
      } catch (error) {
        message.error('Ошибка при загрузке файла');
        onError(error);
      } finally {
        setUploading(false);
      }
    } else {
      // При создании водителя сохраняем файл временно
      const reader = new FileReader();
      reader.onloadend = () => {
        if (type === 'licenseScan') {
          setLicenseScanUrl(reader.result as string);
        } else if (type === 'passportScan') {
          setPassportScanUrl(reader.result as string);
        } else {
          setTempPhotoFile(file);
          setPhotoUrl(reader.result as string);
        }
      };
      reader.readAsDataURL(file);
      onSuccess(file);
      message.success('Файл выбран');
    }
  };

  const handleSubmit = async (values: any) => {
    try {
      const data: CreateDriverRequest | UpdateDriverRequest = {
        ...values,
        birthDate: values.birthDate ? values.birthDate.toISOString() : undefined,
        driverLicenseIssueDate: values.driverLicenseIssueDate.toISOString(),
        driverLicenseExpiryDate: values.driverLicenseExpiryDate.toISOString(),
        passportIssueDate: values.passportIssueDate ? values.passportIssueDate.toISOString() : undefined,
        photo: isEdit ? photoUrl : undefined, // При создании не передаем photoUrl (это base64)
      };

      if (isEdit) {
        await updateDriver.mutateAsync({ id: driver.id, data });
        message.success('Водитель успешно обновлен');
      } else {
        const createdDriver = await createDriver.mutateAsync(data as CreateDriverRequest);

        // Если есть временное фото, загружаем его после создания водителя
        if (tempPhotoFile) {
          try {
            const result = await uploadPhoto.mutateAsync({ id: createdDriver.id, file: tempPhotoFile });
            // Обновляем водителя с фото
            await updateDriver.mutateAsync({
              id: createdDriver.id,
              data: { photo: result.filePath }
            });
          } catch (uploadError) {
            message.warning('Водитель создан, но фото не загружено');
          }
        }

        message.success('Водитель успешно создан');
      }

      onSuccess();
    } catch (error) {
      message.error(isEdit ? 'Ошибка при обновлении водителя' : 'Ошибка при создании водителя');
    }
  };

  return (
    <Form
      form={form}
      layout="vertical"
      onFinish={handleSubmit}
      initialValues={{
        status: 'active',
      }}
    >      {/* Строка 1: Фамилия, Имя, Отчество, Статус */}
      <Row gutter={16}>
        <Col span={6}>
          <Form.Item
            label="Фамилия"
            name="lastName"
            rules={[{ required: true, message: 'Введите фамилию' }]}
          >
            <Input placeholder="Иванов" />
          </Form.Item>
        </Col>
        <Col span={6}>
          <Form.Item
            label="Имя"
            name="firstName"
            rules={[{ required: true, message: 'Введите имя' }]}
          >
            <Input placeholder="Иван" />
          </Form.Item>
        </Col>
        <Col span={6}>
          <Form.Item label="Отчество" name="middleName">
            <Input placeholder="Иванович" />
          </Form.Item>
        </Col>
        <Col span={6}>
          <Form.Item
            label="Статус"
            name="status"
            rules={[{ required: true, message: 'Выберите статус' }]}
          >
            <Select>
              <Option value="active">Активен</Option>
              <Option value="inactive">Неактивен</Option>
              <Option value="blocked">Заблокирован</Option>
            </Select>
          </Form.Item>
        </Col>
      </Row>
      {/* Строка 2: Фото водителя, Телефон, Email, Дата рождения */}
      <Row gutter={16}>
<Col span={6}>
          <Form.Item label="Фото водителя">
            <div style={{ display: 'flex', alignItems: 'flex-start', gap: 16 }}>
              {photoUrl ? (
                <div style={{ position: 'relative' }}>
                  <Image
                    src={photoUrl.startsWith('data:') ? photoUrl : `http://localhost:8080${photoUrl}`}
                    alt="Фото водителя"
                    width={120}
                    height={120}
                    style={{ objectFit: 'cover', borderRadius: 8 }}
                    preview={{
                      mask: 'Просмотр',
                    }}
                  />
                  <div style={{ position: 'absolute', top: -8, right: -8, display: 'flex', gap: 4 }}>
                    <input
                      type="file"
                      accept="image/*"
                      style={{ display: 'none' }}
                      id="photo-upload"
                      onChange={(e) => {
                        const file = e.target.files?.[0];
                        if (file && beforeUpload(file)) {
                          customRequest({ file, onSuccess: () => {}, onError: () => {} });
                        }
                      }}
                    />
                    <Button
                      type="primary"
                      icon={<PlusOutlined />}
                      size="small"
                      onClick={() => document.getElementById('photo-upload')?.click()}
                      disabled={uploading}
                      style={{
                        borderRadius: '50%',
                        width: 28,
                        height: 28,
                        padding: 0,
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center',
                      }}
                    />
                    <Button
                      type="primary"
                      danger
                      icon={<DeleteOutlined />}
                      size="small"
                      onClick={handlePhotoDelete}
                      loading={deletePhoto.isPending}
                      style={{
                        borderRadius: '50%',
                        width: 28,
                        height: 28,
                        padding: 0,
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center',
                      }}
                    />
                  </div>
                </div>
              ) : (
                <div
                  style={{
                    width: 120,
                    height: 120,
                    borderRadius: 8,
                    backgroundColor: '#f0f0f0',
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center',
                    fontSize: 12,
                    color: '#999',
                    border: '2px dashed #d9d9d9',
                    position: 'relative',
                  }}
                >
                  Нет фото
                  <input
                    type="file"
                    accept="image/*"
                    style={{ display: 'none' }}
                    id="photo-upload-empty"
                    onChange={(e) => {
                      const file = e.target.files?.[0];
                      if (file && beforeUpload(file)) {
                        customRequest({ file, onSuccess: () => {}, onError: () => {} });
                      }
                    }}
                  />
                  <Button
                    type="primary"
                    icon={<PlusOutlined />}
                    size="small"
                    onClick={() => document.getElementById('photo-upload-empty')?.click()}
                    disabled={uploading}
                    style={{
                      position: 'absolute',
                      top: -8,
                      right: -8,
                      borderRadius: '50%',
                      width: 28,
                      height: 28,
                      padding: 0,
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                    }}
                  />
                </div>
              )}
            </div>
          </Form.Item>
        </Col>        <Col span={18}>
          <Row gutter={16}>
            <Col span={8}>
              <Form.Item
                label="Телефон"
                name="phone"
                rules={[
                  { required: true, message: 'Введите телефон' },
                  { pattern: /^\+7\d{10}$/, message: 'Формат: +79001234567' },
                ]}
              >
                <Input placeholder="+79001234567" />
              </Form.Item>
            </Col>
            <Col span={8}>
              <Form.Item
                label="Email"
                name="email"
                rules={[{ type: 'email', message: 'Некорректный email' }]}
              >
                <Input placeholder="example@mail.ru" />
              </Form.Item>
            </Col>
            <Col span={8}>
              <Form.Item label="Дата рождения" name="birthDate">
                <DatePicker style={{ width: '100%' }} />
              </Form.Item>
            </Col>
          </Row>
        </Col>
      </Row>

      {/* Строка 3: Скан ВУ, Номер ВУ, Дата выдачи, Срок действия, Стаж */}
      <Row gutter={16}>
<Col span={6}>
          <Form.Item label="Скан водительских прав">
            <div style={{ display: 'flex', alignItems: 'flex-start', gap: 16 }}>
              {licenseScanUrl ? (
                <div style={{ position: 'relative' }}>
                  <Image
                    src={licenseScanUrl.startsWith('data:') ? licenseScanUrl : `http://localhost:8080${licenseScanUrl}`}
                    alt="Скан прав"
                    width={120}
                    height={120}
                    style={{ objectFit: 'cover', borderRadius: 8 }}
                    preview={{ mask: 'Просмотр' }}
                  />
                  <div style={{ position: 'absolute', top: -8, right: -8, display: 'flex', gap: 4 }}>
                    <input
                      type="file"
                      accept="image/*"
                      style={{ display: 'none' }}
                      id="license-upload"
                      onChange={(e) => {
                        const file = e.target.files?.[0];
                        if (file && beforeUpload(file)) {
                          customRequest({ file, onSuccess: () => {}, onError: () => {}, type: 'licenseScan' });
                        }
                      }}
                    />
                    <Button
                      type="primary"
                      icon={<PlusOutlined />}
                      size="small"
                      onClick={() => document.getElementById('license-upload')?.click()}
                      disabled={uploading}
                      style={{
                        borderRadius: '50%',
                        width: 28,
                        height: 28,
                        padding: 0,
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center',
                      }}
                    />
                    <Button
                      type="primary"
                      danger
                      icon={<DeleteOutlined />}
                      size="small"
                      onClick={handleLicenseDelete}
                      loading={deleteLicense.isPending}
                      style={{
                        borderRadius: '50%',
                        width: 28,
                        height: 28,
                        padding: 0,
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center',
                      }}
                    />
                  </div>
                </div>
              ) : (
                <div style={{
                  width: 120, height: 120, borderRadius: 8, backgroundColor: '#f0f0f0',
                  display: 'flex', alignItems: 'center', justifyContent: 'center',
                  fontSize: 12, color: '#999', border: '2px dashed #d9d9d9',
                  position: 'relative'
                }}>
                  Нет скана
                  <input
                    type="file"
                    accept="image/*"
                    style={{ display: 'none' }}
                    id="license-upload-empty"
                    onChange={(e) => {
                      const file = e.target.files?.[0];
                      if (file && beforeUpload(file)) {
                        customRequest({ file, onSuccess: () => {}, onError: () => {}, type: 'licenseScan' });
                      }
                    }}
                  />
                  <Button
                    type="primary"
                    icon={<PlusOutlined />}
                    size="small"
                    onClick={() => document.getElementById('license-upload-empty')?.click()}
                    disabled={uploading}
                    style={{
                      position: 'absolute',
                      top: -8,
                      right: -8,
                      borderRadius: '50%',
                      width: 28,
                      height: 28,
                      padding: 0,
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                    }}
                  />
                </div>
              )}
            </div>
          </Form.Item>
        </Col>        <Col span={18}>
          <Row gutter={16}>
            <Col span={8}>
              <Form.Item
                label="Номер ВУ"
                name="driverLicenseNumber"
                rules={[{ required: true, message: 'Введите номер ВУ' }]}
              >
                <Input placeholder="1234567890" />
              </Form.Item>
            </Col>
            <Col span={8}>
              <Form.Item
                label="Дата выдачи ВУ"
                name="driverLicenseIssueDate"
                rules={[{ required: true, message: 'Выберите дату выдачи' }]}
              >
                <DatePicker style={{ width: '100%' }} />
              </Form.Item>
            </Col>
            <Col span={8}>
              <Form.Item
                label="Срок действия ВУ"
                name="driverLicenseExpiryDate"
                rules={[{ required: true, message: 'Выберите срок действия' }]}
              >
                <DatePicker style={{ width: '100%' }} />
              </Form.Item>
            </Col>
          </Row>        </Col>
      </Row>

      {/* Строка 4: Скан паспорта, Серия, Номер, Дата выдачи */}
      <Row gutter={16}>
<Col span={6}>
          <Form.Item label="Скан паспорта">
            <div style={{ display: 'flex', alignItems: 'flex-start', gap: 16 }}>
              {passportScanUrl ? (
                <div style={{ position: 'relative' }}>
                  <Image
                    src={passportScanUrl.startsWith('data:') ? passportScanUrl : `http://localhost:8080${passportScanUrl}`}
                    alt="Скан паспорта"
                    width={120}
                    height={120}
                    style={{ objectFit: 'cover', borderRadius: 8 }}
                    preview={{ mask: 'Просмотр' }}
                  />
                  <div style={{ position: 'absolute', top: -8, right: -8, display: 'flex', gap: 4 }}>
                    <input
                      type="file"
                      accept="image/*"
                      style={{ display: 'none' }}
                      id="passport-upload"
                      onChange={(e) => {
                        const file = e.target.files?.[0];
                        if (file && beforeUpload(file)) {
                          customRequest({ file, onSuccess: () => {}, onError: () => {}, type: 'passportScan' });
                        }
                      }}
                    />
                    <Button
                      type="primary"
                      icon={<PlusOutlined />}
                      size="small"
                      onClick={() => document.getElementById('passport-upload')?.click()}
                      disabled={uploading}
                      style={{
                        borderRadius: '50%',
                        width: 28,
                        height: 28,
                        padding: 0,
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center',
                      }}
                    />
                    <Button
                      type="primary"
                      danger
                      icon={<DeleteOutlined />}
                      size="small"
                      onClick={handlePassportDelete}
                      loading={deletePassport.isPending}
                      style={{
                        borderRadius: '50%',
                        width: 28,
                        height: 28,
                        padding: 0,
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center',
                      }}
                    />
                  </div>
                </div>
              ) : (
                <div style={{
                  width: 120, height: 120, borderRadius: 8, backgroundColor: '#f0f0f0',
                  display: 'flex', alignItems: 'center', justifyContent: 'center',
                  fontSize: 12, color: '#999', border: '2px dashed #d9d9d9',
                  position: 'relative'
                }}>
                  Нет скана
                  <input
                    type="file"
                    accept="image/*"
                    style={{ display: 'none' }}
                    id="passport-upload-empty"
                    onChange={(e) => {
                      const file = e.target.files?.[0];
                      if (file && beforeUpload(file)) {
                        customRequest({ file, onSuccess: () => {}, onError: () => {}, type: 'passportScan' });
                      }
                    }}
                  />
                  <Button
                    type="primary"
                    icon={<PlusOutlined />}
                    size="small"
                    onClick={() => document.getElementById('passport-upload-empty')?.click()}
                    disabled={uploading}
                    style={{
                      position: 'absolute',
                      top: -8,
                      right: -8,
                      borderRadius: '50%',
                      width: 28,
                      height: 28,
                      padding: 0,
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                    }}
                  />
                </div>
              )}
            </div>
          </Form.Item>
        </Col>
        <Col span={18}>
          <Row gutter={16}>
            <Col span={8}>
              <Form.Item label="Серия паспорта" name="passportSeries">
                <Input placeholder="1234" maxLength={4} />
              </Form.Item>
            </Col>
            <Col span={8}>
              <Form.Item label="Номер паспорта" name="passportNumber">
                <Input placeholder="123456" maxLength={6} />
              </Form.Item>
            </Col>
            <Col span={8}>
              <Form.Item label="Дата выдачи паспорта" name="passportIssueDate">
                <DatePicker style={{ width: '100%' }} />
              </Form.Item>
            </Col>
          </Row>
          <Form.Item label="Адрес" name="address">
            <Input.TextArea rows={2} placeholder="Адрес проживания" />
          </Form.Item>
        </Col>
      </Row>
      <Form.Item>
        <Space>
          <Button
            type="primary"
            htmlType="submit"
            loading={createDriver.isPending || updateDriver.isPending}
          >
            {isEdit ? 'Сохранить' : 'Создать'}
          </Button>
          <Button onClick={onCancel}>Отмена</Button>
        </Space>
      </Form.Item>
    </Form>
  );
};
