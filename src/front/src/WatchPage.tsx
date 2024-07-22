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

const WatchDataSecond: React.FC<WatchDataFirstProps> = ({ setDataState }) => {
    const [isSwitchOn, setIsSwitchOn] = useState(false);
    const [isSecondHand, setIsSecondHand] = useState(true);
    const [isLegal, setIsLegal] = useState(false);
    const [isOriginalBox, setIsOriginalBox] = useState(false);
    const [isReferenceNumber, setIsReferenceNumber] = useState(false);
    const [isSerialNumber, setIsSerialNumber] = useState(false);

    const handleChangeState = () => {
        setDataState(2);
    }

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
        <div className={`${styles.WatchDataChild} ${styles.LessPadding}`}>
            <div className={styles.Title}>Watch Type</div>
            <div className={`${styles.Item} ${styles.SpaceTop}`}>
                <SwitchButton onToggle={handleSecondHand} />
                <div className={styles.Text}>Pocket Watch</div>
            </div>
            <div className={styles.Item}>
                <SwitchButton onToggle={handleLegal} />
                <div className={styles.Text}>Wrist Watch</div>
            </div>
            <div className={styles.Item}>
                <SwitchButton onToggle={handleOriginalBox} />
                <div className={styles.Text}>Others</div>
            </div>
            <div className={`${styles.Item} ${styles.SpaceTop}`}>
                <SwitchButton onToggle={handleOriginalBox} />
                <div className={styles.Text}>Men</div>
            </div>
            <div className={styles.Item}>
                <SwitchButton onToggle={handleOriginalBox} />
                <div className={styles.Text}>Women</div>
            </div>
            <div className={`${styles.Title} ${styles.SpaceTop}`}>Year of manufacture</div>
            <YearSelector/>
            <RoundButton onClick={handleChangeState}/>
        </div>
    );
  }

  const YearSelector: React.FC = () => {
    const currentYear = new Date().getFullYear();
    const startYear = 1800;
    const years: number[] = [];
  
    for (let year = currentYear; year >= startYear; year--) {
      years.push(year);
    }
  
    return (
      <div className={styles.YearSelector}>
        <select>
          {years.map(year => (
            <option key={year} value={year}>{year}</option>
          ))}
        </select>
      </div>
    );
  }

  const DiameterSelector: React.FC = () => {
    const currentDiameter = 40;
    const startDiameter = 0;
    const Diameters: number[] = [];
  
    for (let diameter = currentDiameter; diameter >= startDiameter; diameter--) {
      Diameters.push(diameter);
    }
  
    return (
      <div className={styles.DiameterSelector}>
        <select>
          {Diameters.map(year => (
            <option key={year} value={year}>{year} mm</option>
          ))}
        </select>
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

  const WatchDataThird: React.FC<WatchDataFirstProps> = ({ setDataState }) => {
    const [isSwitchOn, setIsSwitchOn] = useState(false);
    const [isSecondHand, setIsSecondHand] = useState(true);
    const [isLegal, setIsLegal] = useState(false);
    const [isOriginalBox, setIsOriginalBox] = useState(false);
    const [isReferenceNumber, setIsReferenceNumber] = useState(false);
    const [isSerialNumber, setIsSerialNumber] = useState(false);

    const handleChangeState = () => {
        setDataState(4);
    }

    const handleSecondHand = (isOn: boolean) => {
        setIsSecondHand(isOn);
    };

    const handleLegal = (isOn: boolean) => {
        setIsLegal(isOn);
    };

    const handleOriginalBox = (isOn: boolean) => {
        setIsOriginalBox(isOn);
    };

    return (
        <div className={styles.WatchDataChild}>
             <div className={styles.Title}>Details</div>
            <div className={`${styles.Item} ${styles.SpaceTop}`}>
                <SwitchButton onToggle={handleSecondHand} />
                <div className={styles.Text}>Automatic Winding</div>
            </div>
            <div className={styles.Item}>
                <SwitchButton onToggle={handleLegal} />
                <div className={styles.Text}>Manual Winding</div>
            </div>
            <div className={styles.Item}>
                <SwitchButton onToggle={handleOriginalBox} />
                <div className={styles.Text}>Quartz</div>
            </div>
            <div className={`${styles.Title} ${styles.SpaceTop}`}>Diameter</div>
            <div className={styles.SetInLine}>
                <DiameterSelector/>
                <DiameterSelector/>
            </div>
            <RoundButton onClick={handleChangeState}/>

        </div>
    )
}

const WatchData: React.FC = () => {
    const [DataState, setDataState] = useState(0);
    return (
        <div className={styles.WatchData}>
            <div className={styles.Name}>Tag Heuer Carrera</div>
            {DataState === 0 && <WatchDataFirst setDataState={setDataState}/>}
            {DataState === 1 && <WatchDataSecond setDataState={setDataState}/>}
            {DataState === 2 && <WatchDataThird setDataState={setDataState}/>}

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