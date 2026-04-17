import React, { useState, useEffect } from 'react';
import { Modal, Image, Button, Typography } from 'antd';
import { LeftOutlined, RightOutlined, CloseOutlined } from '@ant-design/icons';
const { Text } = Typography;

interface ImageSliderProps {
  images: string[];
  initialIndex?: number;
  visible: boolean;
  onClose: () => void;
}

export const ImageSlider: React.FC<ImageSliderProps> = ({
  images,
  initialIndex = 0,
  visible,
  onClose,
}) => {
  const [currentIndex, setCurrentIndex] = useState(initialIndex);

  useEffect(() => {
    setCurrentIndex(initialIndex);
  }, [initialIndex, visible]);

  const handlePrev = () => {
    setCurrentIndex((prev) => (prev === 0 ? images.length - 1 : prev - 1));
  };

  const handleNext = () => {
    setCurrentIndex((prev) => (prev === images.length - 1 ? 0 : prev + 1));
  };

  const handleKeyDown = (e: KeyboardEvent) => {
    if (e.key === 'ArrowLeft') {
      handlePrev();
    } else if (e.key === 'ArrowRight') {
      handleNext();
    } else if (e.key === 'Escape') {
      onClose();
    }
  };

  useEffect(() => {
    if (visible) {
      window.addEventListener('keydown', handleKeyDown);
      return () => window.removeEventListener('keydown', handleKeyDown);
    }
  }, [visible, handleKeyDown]);
  if (images.length === 0) {
    return null;
  }

  const currentImage = images[currentIndex];

  return (
    <Modal
      open={visible}
      onCancel={onClose}
      footer={null}
      width="90vw"
      style={{ top: 20 }}
      bodyStyle={{ padding: 0 }}
      closeIcon={<CloseOutlined />}
    >      <div
        style={{
          position: 'relative',
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          justifyContent: 'center',
          minHeight: '70vh',
          padding: '20px',
        }}
      >
        {/* Кнопка навигации влево */}
        {images.length > 1 && (
          <Button
            type="text"
            icon={<LeftOutlined />}
            onClick={handlePrev}
            style={{
              position: 'absolute',
              left: 10,
              top: '50%',
              transform: 'translateY(-50%)',
              zIndex: 10,
              fontSize: 24,
              height: 60,
              width: 60,
            }}
          />
        )}

        {/* Изображение */}
        <Image
          src={`http://localhost:8081${currentImage}`}
          alt={`Изображение ${currentIndex + 1}`}
          style={{
            maxWidth: '100%',
            maxHeight: '70vh',
            objectFit: 'contain',
          }}
          preview={false}
        />
        {/* Кнопка навигации вправо */}
        {images.length > 1 && (
          <Button
            type="text"
            icon={<RightOutlined />}
            onClick={handleNext}
            style={{
              position: 'absolute',
              right: 10,
              top: '50%',
              transform: 'translateY(-50%)',
              zIndex: 10,
              fontSize: 24,
              height: 60,
              width: 60,
            }}
          />
        )}

        {/* Индикатор текущего изображения */}
        {images.length > 1 && (
          <div
            style={{
              position: 'absolute',
              bottom: 20,
              left: '50%',
              transform: 'translateX(-50%)',
              backgroundColor: 'rgba(0, 0, 0, 0.7)',
              color: 'white',
              padding: '8px 16px',
              borderRadius: 20,
              zIndex: 10,
            }}
          >
            <Text style={{ color: 'white', fontSize: 14 }}>
              {currentIndex + 1} / {images.length}
            </Text>
          </div>
        )}
      </div>
    </Modal>
  );
};
