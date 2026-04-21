import React, { useEffect, useState } from 'react';
import { Form, Input, InputNumber, Select, Button, Row, Col, message, Typography, Checkbox } from 'antd';
import { useCreateVehicle, useUpdateVehicle, useUploadVehicleImages } from '../../hooks/useVehicles';
import { VehicleImageUpload } from './VehicleImageUpload';
import type { Vehicle, CreateVehicleRequest, UpdateVehicleRequest } from '../../types/vehicle';

const { TextArea } = Input;
const { Text } = Typography;

interface VehicleFormProps {
  vehicle?: Vehicle;
  onSuccess: () => void;
  onCancel: () => void;
}

export const VehicleForm: React.FC<VehicleFormProps> = ({ vehicle, onSuccess, onCancel }) => {
  const [form] = Form.useForm();
  const createVehicle = useCreateVehicle();
  const updateVehicle = useUpdateVehicle();
  const uploadImages = useUploadVehicleImages();
  const isEdit = !!vehicle;

  // Состояние для изображений
  const [imgArray, setImgArray] = useState<string[]>(vehicle?.imgArray || []);
  const [mainImageIndex, setMainImageIndex] = useState<number>(vehicle?.mainImageIndex || 1);
  const [pendingFiles, setPendingFiles] = useState<File[]>([]); // Файлы, которые нужно загрузить
  useEffect(() => {
    if (vehicle) {
      form.setFieldsValue(vehicle);
      setImgArray(vehicle.imgArray || []);
      // С сервера уже приходит 1-based индекс (конвертация в ToResponse)
      setMainImageIndex(vehicle.mainImageIndex || 1);
    }
  }, [vehicle, form]);

  const handleSubmit = async (values: any) => {
    try {
      // Конвертируем полные URL в относительные пути для отправки на сервер
      // Blob URL не отправляем на сервер - они будут загружены отдельно через pendingFiles
      const cleanImgArray = imgArray
        .filter(img => !img.startsWith('blob:')) // Убираем blob URL
        .map(img => {
          if (img.startsWith('http://localhost:8081')) {
            return img.replace('http://localhost:8081', '');
          }
          return img;
        });
      // rostechReg оставляем как boolean
      const processedValues = {
        ...values,
      };
      // Удаляем поле fldMainImageIndex, так как оно не нужно на бэкенде
      delete processedValues.fldMainImageIndex;

      // Создаем объект только с измененными полями
      const data: UpdateVehicleRequest = {};

      // Добавляем только те поля, которые изменились или обязательны
      if (processedValues.garageNumber !== vehicle?.garageNumber) data.garageNumber = processedValues.garageNumber;
      if (processedValues.vin !== vehicle?.vin) data.vin = processedValues.vin;
      if (processedValues.brand !== vehicle?.brand) data.brand = processedValues.brand;
      if (processedValues.machine !== vehicle?.machine) data.machine = processedValues.machine;
      if (processedValues.type !== vehicle?.type) data.type = processedValues.type;
      if (processedValues.height !== vehicle?.height) data.height = processedValues.height;
      if (processedValues.power !== vehicle?.power) data.power = processedValues.power;
      if (processedValues.length !== vehicle?.length) data.length = processedValues.length;
      if (processedValues.width !== vehicle?.width) data.width = processedValues.width;
      if (processedValues.heightTs !== vehicle?.heightTs) data.heightTs = processedValues.heightTs;
      if (processedValues.widthWithSupports !== vehicle?.widthWithSupports) data.widthWithSupports = processedValues.widthWithSupports;
      if (processedValues.mass !== vehicle?.mass) data.mass = processedValues.mass;
      if (processedValues.cradleWidthFolded !== vehicle?.cradleWidthFolded) data.cradleWidthFolded = processedValues.cradleWidthFolded;
      if (processedValues.cradleWidthExtended !== vehicle?.cradleWidthExtended) data.cradleWidthExtended = processedValues.cradleWidthExtended;
      if (processedValues.cradleLengthFolded !== vehicle?.cradleLengthFolded) data.cradleLengthFolded = processedValues.cradleLengthFolded;
      if (processedValues.cradleLengthExtended !== vehicle?.cradleLengthExtended) data.cradleLengthExtended = processedValues.cradleLengthExtended;
      if (processedValues.price5 !== vehicle?.price5) data.price5 = processedValues.price5;
      if (processedValues.price22 !== vehicle?.price22) data.price22 = processedValues.price22;
      if (processedValues.status !== vehicle?.status) data.status = processedValues.status;
      if (processedValues.rostechReg !== vehicle?.rostechReg) data.rostechReg = processedValues.rostechReg;
      if (processedValues.special !== vehicle?.special) data.special = processedValues.special;
      if (processedValues.description !== vehicle?.description) data.description = processedValues.description;

      // Проверяем изменения в изображениях
      const oldImgArray = (vehicle?.imgArray || []).map(img => {
        if (img.startsWith('http://localhost:8081')) {
          return img.replace('http://localhost:8081', '');
        }
        return img;
      });
      const oldMainIndex = vehicle?.mainImageIndex || 1;

      // Сравниваем массивы изображений
      const imagesChanged = JSON.stringify(oldImgArray) !== JSON.stringify(cleanImgArray);
      const mainIndexChanged = oldMainIndex !== mainImageIndex;

      if (imagesChanged) {
        data.imgArray = cleanImgArray;
      }
      if (mainIndexChanged) {
        data.mainImageIndex = mainImageIndex - 1;
      }

      if (isEdit) {
        // При редактировании сначала обновляем данные автовышки
        try {
          await updateVehicle.mutateAsync({ id: vehicle.id, data });

          // Загружаем новые изображения из pendingFiles
          if (pendingFiles.length > 0) {
            await uploadImages.mutateAsync({ id: vehicle.id, files: pendingFiles });
          }

          // Удаление старых изображений происходит через обновление массива imgArray
          // Бэкенд сам обрабатывает разницу между старым и новым массивом
          message.success('Автовышка успешно обновлена');
          onSuccess();
        } catch (updateError: any) {
          console.error('Ошибка обновления:', updateError);
          console.error('Детали ошибки:', updateError.response?.data);
          console.error('Статус:', updateError.response?.status);
          throw updateError;
        }
      } else {
        // При создании сначала создаем автовышку с данными
        const createData: CreateVehicleRequest = {
          ...processedValues,
          imgArray: cleanImgArray,
          mainImageIndex: mainImageIndex - 1,
        };

        const createdVehicle = await createVehicle.mutateAsync(createData);

        // Загружаем новые изображения из pendingFiles
        if (pendingFiles.length > 0 && createdVehicle?.id) {
          await uploadImages.mutateAsync({ id: createdVehicle.id, files: pendingFiles });
        }

        message.success('Автовышка успешно создана');
        onSuccess();
      }
    } catch (error) {
      message.error(isEdit ? 'Ошибка при обновлении автовышки' : 'Ошибка при создании автовышки');
    }
  };

  const handleImagesChange = (newImages: string[], newMainIndex: number, files?: File[]) => {
    setImgArray(newImages);
    setMainImageIndex(newMainIndex);
    if (files) {
      setPendingFiles(files);
    }
  };  return (
    <Form
      form={form}
      layout="vertical"
      onFinish={handleSubmit}
      initialValues={{
        garageNumber: '',
        vin: '',
        height: 0,
        type: undefined,
        power: 0,
        price5: 0,
        price22: 0,
        description: '',
        brand: '',
        machine: '',
        length: 0,
        width: 0,
        heightTs: 0,
        widthWithSupports: 0,
        mass: 0,
        cradleWidthFolded: 0,
        cradleWidthExtended: 0,
        cradleLengthFolded: 0,
        cradleLengthExtended: 0,
        special: '',
        rostechReg: false,
        status: 'active',
      }}
    >
      {/* Загрузка изображений */}
      <VehicleImageUpload
        images={imgArray}
        mainImageIndex={mainImageIndex}
        onChange={handleImagesChange}
      />

      {/* Основная информация */}
      <Text strong style={{ fontSize: 16, display: 'block', marginBottom: 16 }}>Основная информация</Text>

      <Row gutter={16}>
        <Col span={12}>
          <Form.Item
            label="Гаражный номер"
            name="garageNumber"
            rules={[{ required: true, message: 'Введите гаражный номер' }]}
          >
            <Input placeholder="Гаражный номер" />
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item
            label="VIN код"
            name="vin"
            rules={[{ required: true, message: 'Введите VIN код' }]}
          >
            <Input placeholder="VIN код" />
          </Form.Item>
        </Col>
      </Row>

      <Row gutter={16}>
        <Col span={12}>
          <Form.Item
            label="Марка"
            name="brand"
          >
            <Input placeholder="Марка автовышки" />
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item
            label="Модель"
            name="machine"
          >
            <Input placeholder="Модель автовышки" />
          </Form.Item>
        </Col>
      </Row>

      <Row gutter={16}>
        <Col span={8}>
          <Form.Item
            label="Тип автовышки"
            name="type"
          >
            <Select
              placeholder="Выберите тип"
              options={[
                { value: 'Телескопическая', label: 'Телескопическая' },
                { value: 'Телескоп + колено', label: 'Телескоп + колено' },
                { value: 'Телескоп + стрела и рукоять', label: 'Телескоп + стрела и рукоять' },
              ]}
            />
          </Form.Item>
        </Col>
        <Col span={8}>
          <Form.Item
            label="Высота подъема (м)"
            name="height"
            rules={[{ required: true, message: 'Введите высоту подъема' }]}
          >
            <InputNumber min={0} style={{ width: '100%' }} placeholder="Высота подъема" />
          </Form.Item>
        </Col>
        <Col span={8}>
          <Form.Item
            label="Мощность (л.с.)"
            name="power"
          >
            <InputNumber min={0} style={{ width: '100%' }} placeholder="Мощность" />
          </Form.Item>
        </Col>
      </Row>

      {/* Габариты */}
      <Text strong style={{ fontSize: 16, display: 'block', marginBottom: 16, marginTop: 24 }}>Габариты</Text>

      <Row gutter={16}>
        <Col span={8}>
          <Form.Item
            label="Длина (м)"
            name="length"
          >
            <InputNumber min={0} style={{ width: '100%' }} placeholder="Длина" />
          </Form.Item>
        </Col>
        <Col span={8}>
          <Form.Item
            label="Ширина (м)"
            name="width"
          >
            <InputNumber min={0} style={{ width: '100%' }} placeholder="Ширина" />
          </Form.Item>
        </Col>
        <Col span={8}>
          <Form.Item
            label="Высота ТС (м)"
            name="heightTs"
          >
            <InputNumber min={0} style={{ width: '100%' }} placeholder="Высота ТС" />
          </Form.Item>
        </Col>
      </Row>

      <Row gutter={16}>
        <Col span={12}>
          <Form.Item
            label="Ширина с опорами (м)"
            name="widthWithSupports"
          >
            <InputNumber min={0} style={{ width: '100%' }} placeholder="Ширина с опорами" />
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item
            label="Масса (кг)"
            name="mass"
          >
            <InputNumber min={0} style={{ width: '100%' }} placeholder="Масса" />
          </Form.Item>
        </Col>
      </Row>

      {/* Люлька */}
      <Text strong style={{ fontSize: 16, display: 'block', marginBottom: 16, marginTop: 24 }}>Люлька</Text>

      <Row gutter={16}>
        <Col span={12}>
          <Form.Item
            label="Ширина люльки сложенная (м)"
            name="cradleWidthFolded"
          >
            <InputNumber min={0} style={{ width: '100%' }} placeholder="Ширина сложенная" />
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item
            label="Ширина люльки разложенная (м)"
            name="cradleWidthExtended"
          >
            <InputNumber min={0} style={{ width: '100%' }} placeholder="Ширина разложенная" />
          </Form.Item>
        </Col>
      </Row>

      <Row gutter={16}>
        <Col span={12}>
          <Form.Item
            label="Длина люльки сложенная (м)"
            name="cradleLengthFolded"
          >
            <InputNumber min={0} style={{ width: '100%' }} placeholder="Длина сложенная" />
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item
            label="Длина люльки разложенная (м)"
            name="cradleLengthExtended"
          >
            <InputNumber min={0} style={{ width: '100%' }} placeholder="Длина разложенная" />
          </Form.Item>
        </Col>
      </Row>

      {/* Цены и статус */}
      <Text strong style={{ fontSize: 16, display: 'block', marginBottom: 16, marginTop: 24 }}>Цены и статус</Text>

      <Row gutter={16}>
        <Col span={12}>
          <Form.Item
            label="Цена 5% НДС (₽)"
            name="price5"
          >
            <InputNumber min={0} style={{ width: '100%' }} placeholder="Цена 5% НДС" />
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item
            label="Цена 22% НДС (₽)"
            name="price22"
          >
            <InputNumber min={0} style={{ width: '100%' }} placeholder="Цена 22% НДС" />
          </Form.Item>
        </Col>
      </Row>

      <Row gutter={16}>
        <Col span={12}>
          <Form.Item
            label="Статус"
            name="status"
            rules={[{ required: true, message: 'Выберите статус' }]}
          >
            <Select
              options={[
                { value: 'active', label: 'Активна' },
                { value: 'inactive', label: 'Неактивна' },
                { value: 'blocked', label: 'Заблокирована' },
              ]}
            />
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item
            label="Регистрация Ростехнадзора"
            name="rostechReg"
            valuePropName="checked"
          >
            <Checkbox>Зарегистрировано</Checkbox>
          </Form.Item>
        </Col>
      </Row>

      {/* Дополнительная информация */}
      <Text strong style={{ fontSize: 16, display: 'block', marginBottom: 16, marginTop: 24 }}>Дополнительная информация</Text>

      <Form.Item
        label="Особенности"
        name="special"
      >
        <TextArea rows={2} placeholder="Особенности автовышки" />
      </Form.Item>

      <Form.Item
        label="Описание"
        name="description"
      >
        <TextArea rows={4} placeholder="Описание автовышки" />
      </Form.Item>

      <Form.Item style={{ marginTop: 24, marginBottom: 0 }}>
        <Button type="primary" htmlType="submit" loading={createVehicle.isPending || updateVehicle.isPending}>
          {isEdit ? 'Обновить' : 'Создать'}
        </Button>
        <Button onClick={onCancel} style={{ marginLeft: 8 }}>
          Отмена
        </Button>
      </Form.Item>
    </Form>
  );
};