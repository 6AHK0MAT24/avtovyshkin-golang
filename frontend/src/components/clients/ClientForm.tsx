import React from 'react';
import {
  Form,
  Input,
  Select,
  DatePicker,
  Button,
  Space,
  Divider,
  Row,
  Col,
} from 'antd';
import type { CreateClientRequest, UpdateClientRequest, ClientType, Client, ClientStatus } from '../../types/client';
import dayjs, { type Dayjs } from 'dayjs';

const { TextArea } = Input;
const { Option } = Select;

interface ClientFormProps {
  client?: Client;
  clientType: ClientType;
  onSubmit: (data: CreateClientRequest | UpdateClientRequest) => void;
  onCancel: () => void;
  loading?: boolean;
}

interface ClientFormValues {
  // Общие поля
  phone: string;
  email?: string;
  address?: string;
  status: ClientStatus;
  notes?: string;

  // Поля для физических лиц
  firstName?: string;
  lastName?: string;
  middleName?: string;
  birthDate?: Dayjs;
  passportSeries?: string;
  passportNumber?: string;
  passportIssueDate?: Dayjs;
  passportIssuedBy?: string;
  inn?: string;

  // Поля для юридических лиц
  companyName?: string;
  companyLegalName?: string;
  ogrn?: string;
  innLegal?: string;
  kpp?: string;
  legalAddress?: string;
  actualAddress?: string;
  bankName?: string;
  bic?: string;
  accountNumber?: string;
  correspondentAccount?: string;
  directorName?: string;
  directorPosition?: string;
  contactPerson?: string;
  contactPersonPhone?: string;
  contactPersonEmail?: string;
}

export const ClientForm: React.FC<ClientFormProps> = ({
  client,
  clientType,
  onSubmit,
  onCancel,
  loading = false,
}) => {
  const [form] = Form.useForm();

  const handleSubmit = (values: ClientFormValues) => {
    const data: CreateClientRequest | UpdateClientRequest = {
      clientType,
      phone: values.phone,
      email: values.email || undefined,
      address: values.address || undefined,
      status: values.status,
      notes: values.notes || undefined,
    };

    if (clientType === 'individual') {
      Object.assign(data, {
        firstName: values.firstName || undefined,
        lastName: values.lastName || undefined,
        middleName: values.middleName || undefined,
        birthDate: values.birthDate ? values.birthDate.format('YYYY-MM-DD') : undefined,
        passportSeries: values.passportSeries || undefined,
        passportNumber: values.passportNumber || undefined,
        passportIssueDate: values.passportIssueDate ? values.passportIssueDate.format('YYYY-MM-DD') : undefined,
        passportIssuedBy: values.passportIssuedBy || undefined,
        inn: values.inn || undefined,
      });
    } else {
      Object.assign(data, {
        companyName: values.companyName || undefined,
        companyLegalName: values.companyLegalName || undefined,
        ogrn: values.ogrn || undefined,
        innLegal: values.innLegal || undefined,
        kpp: values.kpp || undefined,
        legalAddress: values.legalAddress || undefined,
        actualAddress: values.actualAddress || undefined,
        bankName: values.bankName || undefined,
        bic: values.bic || undefined,
        accountNumber: values.accountNumber || undefined,
        correspondentAccount: values.correspondentAccount || undefined,
        directorName: values.directorName || undefined,
        directorPosition: values.directorPosition || undefined,
        contactPerson: values.contactPerson || undefined,
        contactPersonPhone: values.contactPersonPhone || undefined,
        contactPersonEmail: values.contactPersonEmail || undefined,
      });
    }

    onSubmit(data);
  };

  return (
    <Form
      form={form}
      layout="vertical"
      onFinish={handleSubmit}
      initialValues={{
        status: 'active',
        ...client,
        birthDate: client?.birthDate ? dayjs(client.birthDate) : undefined,
        passportIssueDate: client?.passportIssueDate ? dayjs(client.passportIssueDate) : undefined,
      }}
    >
      {/* Общие поля */}
      <Divider orientationMargin={0}>Общая информация</Divider>
      <Row gutter={16}>
        <Col xs={24} sm={12}>
          <Form.Item
            label="Телефон"
            name="phone"
            rules={[{ required: true, message: 'Введите телефон' }]}
          >
            <Input placeholder="+7 (999) 999-99-99" />
          </Form.Item>
        </Col>
        <Col xs={24} sm={12}>
          <Form.Item label="Email" name="email">
            <Input placeholder="email@example.com" />
          </Form.Item>
        </Col>
        <Col xs={24}>
          <Form.Item label="Адрес" name="address">
            <Input placeholder="Адрес" />
          </Form.Item>
        </Col>
        <Col xs={24} sm={12}>
          <Form.Item label="Статус" name="status">
            <Select>
              <Option value="active">Активен</Option>
              <Option value="inactive">Неактивен</Option>
              <Option value="blocked">Заблокирован</Option>
            </Select>
          </Form.Item>
        </Col>
        <Col xs={24}>
          <Form.Item label="Заметки" name="notes">
            <TextArea rows={3} placeholder="Дополнительная информация" />
          </Form.Item>
        </Col>
      </Row>

      {/* Поля для физических лиц */}
      {clientType === 'individual' && (
        <>
          <Divider orientationMargin={0}>Данные физического лица</Divider>
          <Row gutter={16}>
            <Col xs={24} sm={8}>
              <Form.Item
                label="Фамилия"
                name="lastName"
                rules={[{ required: true, message: 'Введите фамилию' }]}
              >
                <Input placeholder="Иванов" />
              </Form.Item>
            </Col>
            <Col xs={24} sm={8}>
              <Form.Item
                label="Имя"
                name="firstName"
                rules={[{ required: true, message: 'Введите имя' }]}
              >
                <Input placeholder="Иван" />
              </Form.Item>
            </Col>
            <Col xs={24} sm={8}>
              <Form.Item label="Отчество" name="middleName">
                <Input placeholder="Иванович" />
              </Form.Item>
            </Col>
            <Col xs={24} sm={12}>
              <Form.Item label="Дата рождения" name="birthDate">
                <DatePicker style={{ width: '100%' }} format="DD.MM.YYYY" />
              </Form.Item>
            </Col>
            <Col xs={24} sm={6}>
              <Form.Item label="Серия паспорта" name="passportSeries">
                <Input maxLength={4} placeholder="1234" />
              </Form.Item>
            </Col>
            <Col xs={24} sm={6}>
              <Form.Item label="Номер паспорта" name="passportNumber">
                <Input maxLength={6} placeholder="123456" />
              </Form.Item>
            </Col>
            <Col xs={24} sm={12}>
              <Form.Item label="Дата выдачи паспорта" name="passportIssueDate">
                <DatePicker style={{ width: '100%' }} format="DD.MM.YYYY" />
              </Form.Item>
            </Col>
            <Col xs={24}>
              <Form.Item label="Кем выдан паспорт" name="passportIssuedBy">
                <Input placeholder="Отделением МВД по ..." />
              </Form.Item>
            </Col>
            <Col xs={24} sm={12}>
              <Form.Item label="ИНН" name="inn">
                <Input maxLength={12} placeholder="12 цифр" />
              </Form.Item>
            </Col>
          </Row>
        </>
      )}

      {/* Поля для юридических лиц */}
      {clientType === 'legal_entity' && (
        <>
          <Divider orientationMargin={0}>Данные юридического лица</Divider>
          <Row gutter={16}>
            <Col xs={24} sm={12}>
              <Form.Item
                label="Название компании"
                name="companyName"
                rules={[{ required: true, message: 'Введите название компании' }]}
              >
                <Input placeholder="ООО Рога и копыта" />
              </Form.Item>
            </Col>
            <Col xs={24} sm={12}>
              <Form.Item label="Полное юридическое название" name="companyLegalName">
                <Input placeholder="Общество с ограниченной ответственностью Рога и копыта" />
              </Form.Item>
            </Col>
            <Col xs={24} sm={8}>
              <Form.Item
                label="ИНН"
                name="innLegal"
                rules={[{ required: true, message: 'Введите ИНН' }]}
              >
                <Input maxLength={10} placeholder="10 цифр" />
              </Form.Item>
            </Col>
            <Col xs={24} sm={8}>
              <Form.Item label="ОГРН" name="ogrn">
                <Input maxLength={15} placeholder="15 цифр" />
              </Form.Item>
            </Col>
            <Col xs={24} sm={8}>
              <Form.Item label="КПП" name="kpp">
                <Input maxLength={9} placeholder="9 цифр" />
              </Form.Item>
            </Col>
            <Col xs={24}>
              <Form.Item label="Юридический адрес" name="legalAddress">
                <Input placeholder="Юридический адрес" />
              </Form.Item>
            </Col>
            <Col xs={24}>
              <Form.Item label="Фактический адрес" name="actualAddress">
                <Input placeholder="Фактический адрес" />
              </Form.Item>
            </Col>
            <Divider orientationMargin={0}>Банковские реквизиты</Divider>
            <Col xs={24} sm={12}>
              <Form.Item label="Название банка" name="bankName">
                <Input placeholder="ПАО Сбербанк" />
              </Form.Item>
            </Col>
            <Col xs={24} sm={6}>
              <Form.Item label="БИК" name="bic">
                <Input maxLength={9} placeholder="9 цифр" />
              </Form.Item>
            </Col>
            <Col xs={24} sm={6}>
              <Form.Item label="Расчетный счет" name="accountNumber">
                <Input maxLength={20} placeholder="20 цифр" />
              </Form.Item>
            </Col>
            <Col xs={24} sm={12}>
              <Form.Item label="Корр. счет" name="correspondentAccount">
                <Input maxLength={20} placeholder="20 цифр" />
              </Form.Item>
            </Col>
            <Divider orientationMargin={0}>Руководитель и контактное лицо</Divider>
            <Col xs={24} sm={12}>
              <Form.Item label="ФИО директора" name="directorName">
                <Input placeholder="Иванов Иван Иванович" />
              </Form.Item>
            </Col>
            <Col xs={24} sm={12}>
              <Form.Item label="Должность директора" name="directorPosition">
                <Input placeholder="Генеральный директор" />
              </Form.Item>
            </Col>
            <Col xs={24} sm={8}>
              <Form.Item label="Контактное лицо" name="contactPerson">
                <Input placeholder="Петров Петр" />
              </Form.Item>
            </Col>
            <Col xs={24} sm={8}>
              <Form.Item label="Телефон контактного лица" name="contactPersonPhone">
                <Input placeholder="+7 (999) 999-99-99" />
              </Form.Item>
            </Col>
            <Col xs={24} sm={8}>
              <Form.Item label="Email контактного лица" name="contactPersonEmail">
                <Input placeholder="email@example.com" />
              </Form.Item>
            </Col>
          </Row>
        </>
      )}

      <Form.Item>
        <Space>
          <Button type="primary" htmlType="submit" loading={loading}>
            {client ? 'Сохранить' : 'Создать'}
          </Button>
          <Button onClick={onCancel}>
            Отмена
          </Button>
        </Space>
      </Form.Item>
    </Form>
  );
};
