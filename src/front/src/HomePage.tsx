import React from 'react';
import Navbar from './NavBar';
import NavbarCta from './NavCta';
import HomeLanding from './HomeLanding';
import HomeCta from './HomeCta';

    const HomePage: React.FC = () => {
    return (
        <div className='HomePage'>
            <Navbar/>
            <HomeLanding/>
            <HomeCta/>
        </div>
    );
    };

    export default HomePage;