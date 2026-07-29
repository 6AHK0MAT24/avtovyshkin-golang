import React, { useState } from 'react';
import { Card, Tabs, Button, Modal, message } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { useClients, useDeleteClient } from '../hooks/useClients';
import { ClientList } from '../components/clients/ClientList';
import { ClientForm } from '../components/clients/ClientForm';
import type { Client } from '../types/client';

export const ClientsPage = () => {
  const [activeTab, setActiveTab] = useState<'individual' | 'legal'>('individual');
  const [page, setPage] = useState(1);
  const [pageSize] = useState(10);
  const [showModal, setShowModal] = useState(false);
  const [editingClient, setEditingClient] = useState<Client | undefined>(undefined);

  const { data, isLoading } = useClients(activeTab, page, pageSize);
  const deleteClient = useDeleteClient();

  const handleCreate = () => {
    setEditingClient(undefined);
    setShowModal(true);
  };

  const handleEdit = (client: Client) => {
    setEditingClient(client);
    setShowModal(true);
  };

  const handleDelete = async (id: string) => {
    try {
      await deleteClient.mutateAsync(id);
    } catch {
      message.error('Ошибка при удалении клиента');
    }
  };

  const handleFormSubmit = () => {
    setShowModal(false);
    setEditingClient(undefined);
  };

  const handleFormCancel = () => {
    setShowModal(false);
    setEditingClient(undefined);
  };

  const handleTabChange = (key: string) => {
    setActiveTab(key as 'individual' | 'legal');
    setPage(1);
  };

  const tabItems = [
    {
      key: 'individual',
      label: 'Физические лица',
      children: (
        <ClientList
          clients={data?.clients || []}
          loading={isLoading}
          onEdit={handleEdit}
          onDelete={handleDelete}
        />
      ),
    },
    {
      key: 'legal',
      label: 'Юридические лица',
      children: (
        <ClientList
          clients={data?.clients || []}
          loading={isLoading}
          onEdit={handleEdit}
          onDelete={handleDelete}
        />
      ),
    },
  ];

  return (
    <div>
      <div className="mb-4 flex justify-between items-center">
        <h1 className="text-2xl font-bold">Клиенты</h1>
        <Button
          type="primary"
          icon={<PlusOutlined />}
          onClick={handleCreate}
        >
          Добавить {activeTab === 'individual' ? 'физическое лицо' : 'юридическое лицо'}
        </Button>
      </div>

      <Card>
        <Tabs
          activeKey={activeTab}
          onChange={handleTabChange}
          items={tabItems}
        />
      </Card>

      <Modal
        title={editingClient ? 'Редактирование клиента' : 'Создание клиента'}
        open={showModal}
        onCancel={handleFormCancel}
        footer={null}
        width={800}
        destroyOnClose
      >
        <ClientForm
          client={editingClient}
          clientType={activeTab === 'individual' ? 'individual' : 'legal_entity'}
          onSubmit={handleFormSubmit}
          onCancel={handleFormCancel}
        />
      </Modal>
    </div>
  );
};
