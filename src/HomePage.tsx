import React from 'react';
import Navbar from './NavBar';
import NavbarCta from './NavCta';
import HomeLanding from './HomeLanding';

    const HomePage: React.FC = () => {
    return (
        <div className='HomePage'>
            <NavbarCta/>
            <Navbar/>
            <HomeLanding/>
        </div>
    );
    };

    export default HomePage;