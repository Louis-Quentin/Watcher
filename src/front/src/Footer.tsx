import React, { useEffect, useState, useRef } from 'react';
import styles from './css/Footer.module.css';
import exp from 'constants';

const Insta = () => {
    return (
        <div className={styles.instagram}/>
    );
}

const Facebook = () => {
    return (
        <div className={styles.facebook}/>
    );
}

const Tiktok = () => {
    return (
        <div className={styles.tiktok}/>
    );
}

const Yt = () => {
    return (
        <div className={styles.yt}/>
    );
}

const X = () => {
    return (
        <div className={styles.x}/>
    );
}
const Policies = () => {
    return (
        <div className={styles.policies}>
            policies |
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
            <div className={` ${styles.topFooter} ${styles.title}`}>
                <Watches/>
                <AboutUs/>
                <Services/>
                <Company/>
            </div>
            <div className={styles.media}>
                <Insta/>
                <Facebook/>
                <Tiktok/>
                <X/>
                <Yt/>
            </div>
            <div className={styles.bottomFooter}>
                <Policies/>
                <div>@2024 Watcher</div>
            </div>
        </div>
    );
}

export default Footer;