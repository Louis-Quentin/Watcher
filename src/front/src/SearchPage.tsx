import React from 'react';
import Navbar from './NavBar';
import NavbarCta from './NavCta';
import styles from './css/SearchPage.module.css';
import WatchList from './WatchList';
import { Splide } from '@splidejs/splide';


function SearchResultsTitle () {
    return (
        <div className = {styles.SearchResultsTitle}>Results for: ...</div>
    );
}

function SearchPage () {

    return (
        <div>
            <NavbarCta/>
            <Navbar/>
            <SearchResultsTitle/>
        </div>
    )
}

export default SearchPage;