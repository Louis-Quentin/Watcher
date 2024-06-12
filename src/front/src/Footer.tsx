import React, { useEffect, useState, useRef } from 'react';
import styles from './css/Footer.module.css';
import exp from 'constants';

const SocialMedia = () => {
    return (
        <div>
            media
        </div>
    );
}

const Policies = () => {
    return (
        <div>
            policies
        </div>
    );
}

const Watches = () => {
    return (
        <div>
            watches
        </div>
    );
}


const AboutUs = () => {
    return (
        <div>
            about us
        </div>
    );
}


const Services = () => {
    return (
        <div>
            services
        </div>
    );
}


const Company = () => {
    return (
        <div>
            company
        </div>
    );
}


function Footer() {
    return (
        <div className={styles.footer}>
            <div className={styles.topFooter}>
                <Watches/>
                <AboutUs/>
                <Services/>
                <Company/>
            </div>
            <div className={styles.bottomFooter}>
                <SocialMedia/>
                <Policies/>
                <div>@ 2024 Watcher</div>
            </div>
        </div>
    );
}

export default Footer;