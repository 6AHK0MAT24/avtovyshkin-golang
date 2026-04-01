import React, { useState } from 'react';
import { Upload, Button, message, Progress } from 'antd';
import { UploadOutlined } from '@ant-design/icons';
import type { UploadProps } from 'antd';

interface FileUploadProps {
  onUpload: (file: File) => Promise<void>;
  accept?: string;
  maxSize?: number; // in bytes
  buttonText?: string;
}

export const FileUpload: React.FC<FileUploadProps> = ({
  onUpload,
  accept = 'image/*,.pdf',
  maxSize = 10 * 1024 * 1024, // 10MB
  buttonText = 'Загрузить файл',
}) => {
  const [uploading, setUploading] = useState(false);
  const [progress, setProgress] = useState(0);

  const beforeUpload = (file: File) => {
    const isValidSize = file.size <= maxSize;
    if (!isValidSize) {
      message.error(`Размер файла не должен превышать ${maxSize / 1024 / 1024}MB`);
      return false;
    }
    return true;
  };

  const handleUpload = async (file: File) => {
    setUploading(true);
    setProgress(0);

    try {
      // Simulate progress
      const progressInterval = setInterval(() => {
        setProgress((prev) => {
          if (prev >= 90) {
            clearInterval(progressInterval);
            return 90;
          }
          return prev + 10;
        });
      }, 200);

      await onUpload(file);

      clearInterval(progressInterval);
      setProgress(100);
      message.success('Файл успешно загружен');

      setTimeout(() => {
        setProgress(0);
        setUploading(false);
      }, 1000);
    } catch (error) {
      message.error('Ошибка при загрузке файла');
      setProgress(0);
      setUploading(false);
    }

    return false; // Prevent automatic upload
  };

  const uploadProps: UploadProps = {
    beforeUpload,
    customRequest: ({ file }) => handleUpload(file as File),
    showUploadList: false,
    accept,
  };

  return (
    <div>
      <Upload {...uploadProps}>
        <Button icon={<UploadOutlined />} loading={uploading} disabled={uploading}>
          {buttonText}
        </Button>
      </Upload>
      {uploading && <Progress percent={progress} size="small" style={{ marginTop: 8 }} />}
    </div>
  );
};
