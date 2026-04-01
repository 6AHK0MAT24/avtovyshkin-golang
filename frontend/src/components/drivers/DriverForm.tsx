import React, { useEffect, useState } from 'react';
import { Form, Input, InputNumber, DatePicker, Select, Button, Space, message, Upload, Image } from 'antd';
import { PlusOutlined, DeleteOutlined } from '@ant-design/icons';
import dayjs from 'dayjs';
import { useCreateDriver, useUpdateDriver, useUploadDriverPhoto, useDeleteDriverPhoto } from '../../hooks/useDrivers';
import type { Driver, CreateDriverRequest, UpdateDriverRequest } from '../../types/driver';
import type { UploadFile } from 'antd/es/upload/interface';
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
  const [photoUrl, setPhotoUrl] = useState<string | undefined>(driver?.photo);
  const [fileList, setFileList] = useState<UploadFile[]>([]);
  const [uploading, setUploading] = useState(false);
  const [tempPhotoFile, setTempPhotoFile] = useState<File | null>(null);

  const isEdit = !!driver;

  useEffect(() => {    if (driver) {
      form.setFieldsValue({
        ...driver,
        birthDate: driver.birthDate ? dayjs(driver.birthDate) : null,
        driverLicenseIssueDate: dayjs(driver.driverLicenseIssueDate),
        driverLicenseExpiryDate: dayjs(driver.driverLicenseExpiryDate),
        passportIssueDate: driver.passportIssueDate ? dayjs(driver.passportIssueDate) : null,
      });
      setPhotoUrl(driver.photo);
    }
  }, [driver, form]);
  const handlePhotoChange = (info: any) => {
    const { file } = info;

    if (file.status === 'done') {
      const reader = new FileReader();
      reader.onloadend = () => {
        setPhotoUrl(reader.result as string);
      };
      reader.readAsDataURL(file.originFileObj);
    }
  };

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
    setFileList([]);
    setTempPhotoFile(null);
    form.setFieldValue('photo', undefined);
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
    const { file, onSuccess, onError } = options;

    if (isEdit && driver) {
      try {
        setUploading(true);
        const result = await uploadPhoto.mutateAsync({ id: driver.id, file });
        setPhotoUrl(result.filePath);
        form.setFieldValue('photo', result.filePath);
        onSuccess(result);
        message.success('Фото успешно загружено');
      } catch (error) {
        message.error('Ошибка при загрузке фото');
        onError(error);
      } finally {
        setUploading(false);
      }
    } else {
      // При создании водителя сохраняем файл временно
      setTempPhotoFile(file);
      const reader = new FileReader();
      reader.onloadend = () => {
        setPhotoUrl(reader.result as string);
      };
      reader.readAsDataURL(file);
      onSuccess(file);
      message.success('Фото выбрано');
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
  };  return (
    <Form
      form={form}
      layout="vertical"
      onFinish={handleSubmit}
      initialValues={{
        status: 'active',
        experienceYears: 0,
      }}
    >
      <Form.Item
        label="Фамилия"
        name="lastName"
        rules={[{ required: true, message: 'Введите фамилию' }]}
      >
        <Input placeholder="Иванов" />
      </Form.Item>

      <Form.Item
        label="Имя"
        name="firstName"
        rules={[{ required: true, message: 'Введите имя' }]}
      >
        <Input placeholder="Иван" />
      </Form.Item>

      <Form.Item label="Отчество" name="middleName">
        <Input placeholder="Иванович" />
      </Form.Item>

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
              <Button
                type="primary"
                danger
                icon={<DeleteOutlined />}
                size="small"
                onClick={handlePhotoDelete}
                loading={deletePhoto.isPending}
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
              }}
            >
              Нет фото
            </div>
          )}
          <Upload
            listType="picture-card"
            fileList={fileList}
            onChange={handlePhotoChange}
            beforeUpload={beforeUpload}
            customRequest={customRequest}
            maxCount={1}
            accept="image/*"
            showUploadList={false}
            disabled={uploading}
          >
            <div>
              <PlusOutlined />
              <div style={{ marginTop: 8 }}>
                {uploading ? 'Загрузка...' : 'Загрузить'}
              </div>
            </div>
          </Upload>        </div>
      </Form.Item>
      <Form.Item        label="Телефон"
        name="phone"
        rules={[
          { required: true, message: 'Введите телефон' },
          { pattern: /^\+7\d{10}$/, message: 'Формат: +79001234567' },
        ]}
      >
        <Input placeholder="+79001234567" />
      </Form.Item>

      <Form.Item
        label="Email"
        name="email"
        rules={[{ type: 'email', message: 'Некорректный email' }]}
      >
        <Input placeholder="example@mail.ru" />
      </Form.Item>

      <Form.Item label="Дата рождения" name="birthDate">
        <DatePicker style={{ width: '100%' }} />
      </Form.Item>

      <Form.Item
        label="Номер водительского удостоверения"
        name="driverLicenseNumber"
        rules={[{ required: true, message: 'Введите номер ВУ' }]}
      >
        <Input placeholder="1234567890" />
      </Form.Item>

      <Form.Item
        label="Дата выдачи ВУ"
        name="driverLicenseIssueDate"
        rules={[{ required: true, message: 'Выберите дату выдачи' }]}
      >
        <DatePicker style={{ width: '100%' }} />
      </Form.Item>

      <Form.Item
        label="Срок действия ВУ"
        name="driverLicenseExpiryDate"
        rules={[{ required: true, message: 'Выберите срок действия' }]}
      >
        <DatePicker style={{ width: '100%' }} />
      </Form.Item>

      <Form.Item label="Серия паспорта" name="passportSeries">
        <Input placeholder="1234" maxLength={4} />
      </Form.Item>

      <Form.Item label="Номер паспорта" name="passportNumber">
        <Input placeholder="123456" maxLength={6} />
      </Form.Item>

      <Form.Item label="Дата выдачи паспорта" name="passportIssueDate">
        <DatePicker style={{ width: '100%' }} />
      </Form.Item>

      <Form.Item label="Адрес" name="address">
        <Input.TextArea rows={3} placeholder="Адрес проживания" />
      </Form.Item>

      <Form.Item
        label="Стаж вождения (лет)"
        name="experienceYears"
        rules={[{ required: true, message: 'Укажите стаж' }]}
      >
        <InputNumber min={0} max={50} style={{ width: '100%' }} />
      </Form.Item>

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
