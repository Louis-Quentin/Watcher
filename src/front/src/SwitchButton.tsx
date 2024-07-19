import React, { useState } from 'react';
import styles from './css/SwitchButton.module.css';  // Adjust the path as necessary

interface SwitchButtonProps {
  onToggle: (isOn: boolean) => void;
}

const SwitchButton: React.FC<SwitchButtonProps> = ({ onToggle }) => {
  const [isOn, setIsOn] = useState(false);

  const toggleSwitch = () => {
    const newIsOn = !isOn;
    setIsOn(newIsOn);
    onToggle(newIsOn); // Call the onToggle function with the new state
  };

  return (
    <label className={styles.switch}>
      <input type="checkbox" checked={isOn} onChange={toggleSwitch} />
      <span className={styles.slider}></span>
    </label>
  );
};

export default SwitchButton;
