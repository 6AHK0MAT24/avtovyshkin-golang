import React, { useState, useRef } from 'react';
import { Upload, message, Button, Typography, Space } from 'antd';
import { InboxOutlined, DeleteOutlined, StarOutlined, StarFilled } from '@ant-design/icons';
import type { UploadProps } from 'antd';
const { Dragger } = Upload;
const { Text } = Typography;

interface VehicleImageUploadProps {
  images?: string[];
  mainImageIndex?: number;
  onChange?: (images: string[], mainImageIndex: number) => void;
  maxCount?: number;
  maxSize?: number; // в байтах
}

interface ImagePreview {
  url: string;
  file?: File;
}

export const VehicleImageUpload: React.FC<VehicleImageUploadProps> = ({
  images = [],
  mainImageIndex = 0,
  onChange,
  maxCount = 20,
  maxSize = 10 * 1024 * 1024, // 10MB
}) => {
  const [previews, setPreviews] = useState<ImagePreview[]>(
    images.map((url) => ({ url }))
  );
  const [currentMainIndex, setCurrentMainIndex] = useState(mainImageIndex);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const allowedTypes = ['image/jpeg', 'image/jpg', 'image/png', 'image/gif', 'image/webp'];

  const validateFile = (file: File): boolean => {
    // Проверка типа файла
    if (!allowedTypes.includes(file.type)) {
      message.error('Допустимы только изображения (jpg, jpeg, png, gif, webp)');
      return false;
    }

    // Проверка размера файла
    if (file.size > maxSize) {
      message.error('Размер файла не должен превышать 10MB');
      return false;
    }

    // Проверка количества файлов
    if (previews.length >= maxCount) {
      message.error(`Максимум ${maxCount} изображений`);
      return false;
    }

    return true;
  };

  const handleFileSelect = (files: FileList | null) => {
    if (!files) return;

    const validFiles: File[] = [];
    const newPreviews: ImagePreview[] = [...previews];

    for (let i = 0; i < files.length; i++) {
      const file = files[i];

      if (validateFile(file)) {
        validFiles.push(file);
        const url = URL.createObjectURL(file);
        newPreviews.push({ url, file });
      }
    }

    if (validFiles.length > 0) {
      setPreviews(newPreviews);
      handleChange(newPreviews, currentMainIndex);
    }
  };

  const handleDrop: UploadProps['customRequest'] = (options) => {
    const { file } = options;
    if (file instanceof File) {
      handleFileSelect([file] as unknown as FileList);
    }
  };

  const handleRemove = (index: number) => {
    const newPreviews = previews.filter((_, i) => i !== index);
    let newMainIndex = currentMainIndex;

    // Если удалили основное изображение, устанавливаем новое основное
    if (index === currentMainIndex) {
      newMainIndex = newPreviews.length > 0 ? 0 : 0;
    } else if (index < currentMainIndex) {
      // Если удалили изображение до основного, сдвигаем индекс
      newMainIndex = currentMainIndex - 1;
    }

    setPreviews(newPreviews);
    setCurrentMainIndex(newMainIndex);
    handleChange(newPreviews, newMainIndex);
  };

  const handleSetMain = (index: number) => {
    setCurrentMainIndex(index);
    handleChange(previews, index);
  };

  const handleChange = (newPreviews: ImagePreview[], newMainIndex: number) => {
    if (onChange) {
      const imageUrls = newPreviews.map((p) => p.url);
      onChange(imageUrls, newMainIndex);
    }
  };
  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    handleFileSelect(e.target.files);
    // Сбрасываем input, чтобы можно было выбрать те же файлы снова
    if (fileInputRef.current) {
      fileInputRef.current.value = '';
    }
  };

  return (
    <div>
      <Space direction="vertical" style={{ width: '100%' }} size="middle">
        {/* Drag & Drop зона */}
        {previews.length < maxCount && (
          <Dragger
            name="images"
            multiple
            accept={allowedTypes.join(',')}
            showUploadList={false}
            customRequest={handleDrop}
            beforeUpload={validateFile}
            style={{ padding: '20px' }}
          >
            <p className="ant-upload-drag-icon">
              <InboxOutlined style={{ fontSize: 48, color: '#1890ff' }} />
            </p>
            <p className="ant-upload-text">
              Перетащите изображения сюда или нажмите для выбора
            </p>
            <p className="ant-upload-hint">
              Поддержка: jpg, jpeg, png, gif, webp (макс. 10MB)
            </p>
          </Dragger>
        )}

        {/* Счётчик изображений */}
        <Text type="secondary">
          {previews.length} / {maxCount} изображений
        </Text>

        {/* Grid предпросмотра изображений */}
        {previews.length > 0 && (
          <div
            style={{
              display: 'grid',
              gridTemplateColumns: 'repeat(auto-fill, minmax(150px, 1fr))',
              gap: '16px',
            }}
          >
            {previews.map((preview, index) => (
              <div
                key={index}
                style={{
                  position: 'relative',
                  aspectRatio: '1',
                  borderRadius: '8px',
                  overflow: 'hidden',
                  border: currentMainIndex === index ? '3px solid #1890ff' : '2px solid #d9d9d9',
                  boxShadow: currentMainIndex === index ? '0 0 10px rgba(24, 144, 255, 0.3)' : 'none',
                }}
              >
                {/* Изображение */}
                <img
                  src={preview.url}
                  alt={`Изображение ${index + 1}`}
                  style={{
                    width: '100%',
                    height: '100%',
                    objectFit: 'cover',
                  }}
                />

                {/* Кнопка удаления */}
                <Button
                  type="primary"
                  danger
                  shape="circle"
                  icon={<DeleteOutlined />}
                  size="small"
                  onClick={() => handleRemove(index)}
                  style={{
                    position: 'absolute',
                    top: 4,
                    right: 4,
                    zIndex: 2,
                  }}
                />

                {/* Кнопка выбора основного изображения */}
                <Button
                  type={currentMainIndex === index ? 'primary' : 'default'}
                  shape="circle"
                  icon={currentMainIndex === index ? <StarFilled /> : <StarOutlined />}
                  size="small"
                  onClick={() => handleSetMain(index)}
                  style={{
                    position: 'absolute',
                    top: 4,
                    left: 4,
                    zIndex: 2,
                  }}
                  title={currentMainIndex === index ? 'Основное изображение' : 'Сделать основным'}
                />

                {/* Индикатор основного изображения */}
                {currentMainIndex === index && (
                  <div
                    style={{
                      position: 'absolute',
                      bottom: 0,
                      left: 0,
                      right: 0,
                      backgroundColor: 'rgba(24, 144, 255, 0.9)',
                      color: 'white',
                      textAlign: 'center',
                      padding: '4px 8px',
                      fontSize: '12px',
                      fontWeight: 'bold',
                    }}
                  >
                    Основное
                  </div>
                )}
              </div>
            ))}
          </div>
        )}

        {/* Скрытый input для выбора файлов */}
        <input
          ref={fileInputRef}
          type="file"
          multiple
          accept={allowedTypes.join(',')}
          style={{ display: 'none' }}
          onChange={handleInputChange}
        />
      </Space>
    </div>
  );
};
