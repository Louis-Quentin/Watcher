import React from 'react';
import Navbar from './NavBar';
import NavbarCta from './NavCta';
import HomeLanding from './HomeLanding';
import HomeCta from './HomeCta';
import MapCta from './MapCta';
import Footer from './Footer'

    const HomePage: React.FC = () => {
    return (
        <div className='HomePage'>
            <Navbar/>
            <HomeLanding/>
            <HomeCta/>
            <MapCta/>
            <Footer/>
        </div>
    );
    };

    export default HomePage;