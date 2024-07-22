import React from 'react';
import Navbar from './NavBar';
import Footer from './Footer';
import styles from './css/RetailerPage.module.css'

const RetailerPage: React.FC = (retailer: any) => {
    return (
        <div>
            <Navbar/>
            <div className={styles.Retailer}>
                <div className={styles.Name}>{retailer.name}zzz</div>
            </div>
            <Footer/>
        </div>
    );
}

export default RetailerPage;
