import React, { useEffect } from 'react';
import { Form, Input, InputNumber, Select, Button, Row, Col, message } from 'antd';
import { useCreateVehicle, useUpdateVehicle } from '../../hooks/useVehicles';import type { Vehicle, CreateVehicleRequest, UpdateVehicleRequest } from '../../types/vehicle';

const { Option } = Select;
const { TextArea } = Input;

interface VehicleFormProps {
  vehicle?: Vehicle;
  onSuccess: () => void;
  onCancel: () => void;
}

export const VehicleForm: React.FC<VehicleFormProps> = ({ vehicle, onSuccess, onCancel }) => {
  const [form] = Form.useForm();
  const createVehicle = useCreateVehicle();
  const updateVehicle = useUpdateVehicle();
  const isEdit = !!vehicle;

  useEffect(() => {
    if (vehicle) {
      form.setFieldsValue(vehicle);
    }
  }, [vehicle, form]);

  const handleSubmit = async (values: any) => {
    try {
      const data: CreateVehicleRequest | UpdateVehicleRequest = {
        ...values,
      };

      if (isEdit) {
        await updateVehicle.mutateAsync({ id: vehicle.id, data });
        message.success('Автовышка успешно обновлена');
      } else {
        await createVehicle.mutateAsync(data as CreateVehicleRequest);
        message.success('Автовышка успешно создана');
      }

      onSuccess();
    } catch (error) {
      message.error(isEdit ? 'Ошибка при обновлении автовышки' : 'Ошибка при создании автовышки');
    }
  };

  return (
    <Form
      form={form}
      layout="vertical"
      onFinish={handleSubmit}
      initialValues={{
        status: 'active',
        mainImageIndex: 0,
      }}
    >
      {/* Строка 1: Гаражный номер, VIN, Высота, Статус */}
      <Row gutter={16}>
        <Col span={6}>
          <Form.Item
            label="Гаражный номер"
            name="garageNumber"
            rules={[{ required: true, message: 'Введите гаражный номер' }]}
          >
            <Input placeholder="А001" />
          </Form.Item>
        </Col>
        <Col span={6}>
          <Form.Item
            label="VIN"
            name="vin"
            rules={[
              { required: true, message: 'Введите VIN' },
              { len: 17, message: 'VIN должен содержать 17 символов' },
            ]}
          >
            <Input placeholder="WBA12345678901234" />
          </Form.Item>
        </Col>
        <Col span={6}>
          <Form.Item
            label="Высота (м)"
            name="height"
            rules={[
              { required: true, message: 'Введите высоту' },
              { type: 'number', min: 0.1, message: 'Высота должна быть больше 0' },
            ]}
          >
            <InputNumber
              placeholder="22"
              min={0.1}
              step={0.1}
              style={{ width: '100%' }}
            />
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

      {/* Строка 2: Бренд, Тип, Мощность, РостехРег */}
      <Row gutter={16}>
        <Col span={6}>
          <Form.Item label="Бренд" name="brand">
            <Input placeholder="JLG" />
          </Form.Item>
        </Col>
        <Col span={6}>
          <Form.Item
            label="Тип"
            name="type"
          >
            <Select allowClear placeholder="Выберите тип">
              <Option value="Телескопическая">Телескопическая</Option>
              <Option value="Телескоп + колено">Телескоп + колено</Option>
              <Option value="Телескоп + стрела и рукоять">Телескоп + стрела и рукоять</Option>
            </Select>
          </Form.Item>
        </Col>
        <Col span={6}>
          <Form.Item label="Мощность (кВт)" name="power">
            <InputNumber
              placeholder="50"
              min={0}
              step={0.1}
              style={{ width: '100%' }}
            />
          </Form.Item>
        </Col>
        <Col span={6}>
          <Form.Item label="РостехРег" name="rostechReg">
            <Input placeholder="РР-12345" />
          </Form.Item>
        </Col>
      </Row>

      {/* Строка 3: Длина, Ширина, Высота ТС, Ширина с опорами */}
      <Row gutter={16}>
        <Col span={6}>
          <Form.Item label="Длина (м)" name="length">
            <InputNumber
              placeholder="8"
              min={0}
              step={0.1}
              style={{ width: '100%' }}
            />
          </Form.Item>
        </Col>
        <Col span={6}>
          <Form.Item label="Ширина (м)" name="width">
            <InputNumber
              placeholder="2.5"
              min={0}
              step={0.1}
              style={{ width: '100%' }}
            />
          </Form.Item>
        </Col>
        <Col span={6}>
          <Form.Item label="Высота ТС (м)" name="heightTs">
            <InputNumber
              placeholder="3"
              min={0}
              step={0.1}
              style={{ width: '100%' }}
            />
          </Form.Item>
        </Col>
        <Col span={6}>
          <Form.Item label="Ширина с опорами (м)" name="widthWithSupports">
            <InputNumber
              placeholder="4"
              min={0}
              step={0.1}
              style={{ width: '100%' }}
            />
          </Form.Item>
        </Col>
      </Row>

      {/* Строка 4: Масса, Ширина люльки (сложена), Ширина люльки (разложена) */}
      <Row gutter={16}>
        <Col span={6}>
          <Form.Item label="Масса (кг)" name="mass">
            <InputNumber
              placeholder="8000"
              min={0}
              step={1}
              style={{ width: '100%' }}
            />
          </Form.Item>
        </Col>
        <Col span={6}>
          <Form.Item label="Ширина люльки (сложена, м)" name="cradleWidthFolded">
            <InputNumber
              placeholder="1.2"
              min={0}
              step={0.1}
              style={{ width: '100%' }}
            />
          </Form.Item>
        </Col>
        <Col span={6}>
          <Form.Item label="Ширина люльки (разложена, м)" name="cradleWidthExtended">
            <InputNumber
              placeholder="2.5"
              min={0}
              step={0.1}
              style={{ width: '100%' }}
            />
          </Form.Item>
        </Col>
        <Col span={6}>
          <Form.Item label="Машина" name="machine">
            <Input placeholder="Модель машины" />
          </Form.Item>
        </Col>
      </Row>

      {/* Строка 5: Длина люльки (сложена), Длина люльки (разложена) */}
      <Row gutter={16}>
        <Col span={6}>
          <Form.Item label="Длина люльки (сложена, м)" name="cradleLengthFolded">
            <InputNumber
              placeholder="1.5"
              min={0}
              step={0.1}
              style={{ width: '100%' }}
            />
          </Form.Item>
        </Col>
        <Col span={6}>
          <Form.Item label="Длина люльки (разложена, м)" name="cradleLengthExtended">
            <InputNumber
              placeholder="3"
              min={0}
              step={0.1}
              style={{ width: '100%' }}
            />
          </Form.Item>
        </Col>
      </Row>

      {/* Строка 6: Цена (5м), Цена (22м) */}
      <Row gutter={16}>
        <Col span={6}>
          <Form.Item label="Цена (5м, ₽)" name="price5">
            <InputNumber
              placeholder="5000"
              min={0}
              step={100}
              style={{ width: '100%' }}
            />
          </Form.Item>
        </Col>
        <Col span={6}>
          <Form.Item label="Цена (22м, ₽)" name="price22">
            <InputNumber
              placeholder="15000"
              min={0}
              step={100}
              style={{ width: '100%' }}
            />
          </Form.Item>
        </Col>
      </Row>

      {/* Строка 7: Описание, Особые отметки */}
      <Row gutter={16}>
        <Col span={12}>
          <Form.Item label="Описание" name="description">
            <TextArea
              placeholder="Описание автовышки"
              rows={3}
            />
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item label="Особые отметки" name="special">
            <TextArea
              placeholder="Особые отметки"
              rows={3}
            />
          </Form.Item>
        </Col>
      </Row>

      {/* Кнопки действий */}
      <Row gutter={16} style={{ marginTop: 24 }}>
        <Col>
          <Button type="primary" htmlType="submit" loading={createVehicle.isPending || updateVehicle.isPending}>
            {isEdit ? 'Сохранить' : 'Создать'}
          </Button>
        </Col>
        <Col>
          <Button onClick={onCancel}>
            Отмена
          </Button>
        </Col>
      </Row>
    </Form>
  );
};
