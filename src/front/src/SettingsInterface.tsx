import React, { useState } from 'react';
import styles from './css/SettingsInterface.module.css';

interface SettingsInterfaceProps {
  onClose: () => void;
}

const SettingsInterface: React.FC<SettingsInterfaceProps> = ({ onClose }) => {
  const [rangeValue, setRangeValue] = useState(50);

  return (
    <div className={styles.container}>
      <button className={styles.closeButton} onClick={onClose}>
        X
      </button>
      <label>
        Range: {rangeValue} km
        <input
          type="range"
          min="0"
          max="100"
          value={rangeValue}
          onChange={(e) => setRangeValue(Number(e.target.value))}
          className={styles.rangeSlider}
        />
      </label>
    </div>
  );
};

export default SettingsInterface;
