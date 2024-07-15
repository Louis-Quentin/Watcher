import React from 'react';
import Navbar from './NavBar';
import Footer from './Footer'
import BasicCarousel from './BasicCarousel';
import SearchBar from './SearchBar';
import styles from './css/HomeLanding.module.css';

import image1 from './css/img/store_1.jpg';
import image2 from './css/img/store_2.jpg';
import image3 from './css/img/store_3.jpg';
import MapComponent from './Map';


    const StoresPage: React.FC = () => {

        const images = [
            image1,
            image2,
            image3
        ];
    return (
        <div className={styles.HomePage}>
            <Navbar/>
            <SearchBar/>
            <MapComponent/>
            <Footer/>
        </div>
    );
    };

    export default StoresPage;