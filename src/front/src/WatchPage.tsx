import React, { useState } from 'react';
import Navbar from './NavBar';
import Footer from './Footer'
import styles from './css/WatchPage.module.css';
import SwitchButton from './SwitchButton';

interface WatchDataFirstProps {
    setDataState: React.Dispatch<React.SetStateAction<number>>;
  }

interface RoundButtonProps {
    onClick: () => void;
  }

  const WatchDataFirst: React.FC<WatchDataFirstProps> = ({ setDataState }) => {
    const [isSwitchOn, setIsSwitchOn] = useState(false);
    const [isSecondHand, setIsSecondHand] = useState(true);
    const [isLegal, setIsLegal] = useState(false);
    const [isOriginalBox, setIsOriginalBox] = useState(false);
    const [isReferenceNumber, setIsReferenceNumber] = useState(false);
    const [isSerialNumber, setIsSerialNumber] = useState(false);

    const handleChangeState = () => {
        setDataState(1);
    }
  
    const handleToggle = (isOn: boolean) => {
      setIsSwitchOn(isOn);
    };
  
    const handleSecondHand = (isOn: boolean) => {
        setIsSecondHand(isOn);
    };

    const handleLegal = (isOn: boolean) => {
        setIsLegal(isOn);
    };

    const handleOriginalBox = (isOn: boolean) => {
        setIsOriginalBox(isOn);
    };

    const handleReferenceNumber = (isOn: boolean) => {
        setIsReferenceNumber(isOn);
    };

    const handleSerialNumber = (isOn: boolean) => {
        setIsSerialNumber(isOn);
    };
    return (
      <div className={styles.WatchDataChild}>
        <div className={styles.Item}>
          <SwitchButton onToggle={handleToggle} />
          <div className={styles.Text}>Brand New</div>
        </div>
        <div className={`${styles.SecondHand} ${!isSwitchOn ? styles.Overlay : styles.Hidden}`}>
            <div className={styles.Item}>
                <SwitchButton onToggle={handleSecondHand} />
                <div className={styles.Text}>Second Hand</div>
            </div>
            <div className={styles.Item}>
                <SwitchButton onToggle={handleLegal} />
                <div className={styles.Text}>Legal Papers</div>
            </div>
            <div className={styles.Item}>
                <SwitchButton onToggle={handleOriginalBox} />
                <div className={styles.Text}>Original Box</div>
            </div>
            <div className={styles.Item}>
                <SwitchButton onToggle={handleOriginalBox} />
                <div className={styles.Text}>Reference Number</div>
            </div>
            <div className={styles.Item}>
                <SwitchButton onToggle={handleOriginalBox} />
                <div className={styles.Text}>Serial Number</div>
            </div>
        </div>
        <RoundButton onClick={handleChangeState}/>
      </div>
    );
  };

const WatchDataSecond: React.FC = () => {
    const [isSwitchOn, setIsSwitchOn] = useState(false);

    return (
        <div>

        </div>
    );
  }

const RoundButton: React.FC<RoundButtonProps> = ({ onClick }) => {
    return (
      <button className={styles.RoundButton} onClick={onClick}>
        Next &rarr; {/* Right arrow character */}
      </button>
    );
  }

const WatchData: React.FC = () => {
    const [DataState, setDataState] = useState(0);
    return (
        <div className={styles.WatchData}>
            <div className={styles.Name}>Tag Heuer Carrera</div>
            {DataState === 0 && <WatchDataFirst setDataState={setDataState}/>}
        </div>
    );
}

const WatchCard: React.FC = () => {
    return (
        <div className={styles.WatchCard}>
            <img className={styles.WatchImg} src={require('./css/img/tag_heuer_test.jpg')}></img>
            <WatchData/>
        </div>
    );
}

const WatchPage: React.FC = () => {
    return (
        <div>
            <Navbar/>
            <WatchCard/>
            <Footer/>
        </div>
    );
}

export default WatchPage;