import React from 'react';
import styles from './css/WatchList.module.css';
import test from './watchs_sample.json';
import { Watch } from './DataStructure.js';
import { cp } from 'fs';

async function fetchRecommendations(): Promise<Watch[]> {
    try {
      const response = await fetch("http://localhost:8080/recommendation");
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      return data as Watch[];
    } catch (error) {
      console.error("Error fetching recommendations:", error);
      return [];
    }
  }

/*const DisplayWatchList = ({ watch }: { watch: Watch }) => {
    return(
        <div className= {styles.comp}>
            <div className= {styles.card}>
                <div className= {styles.textbox}>{watch.Range}</div>
                <div className= {styles.textbox}>{watch.Name}</div>
                <div className= {styles.textbox}>{watch.Url}</div>
                <div className= {styles.textbox}>{watch.Available}</div>
                <div className= {styles.textbox}>{watch.Brand}</div>
            </div>
        </div>
    );
}*/

async function WatchList() {
    const data = await fetchRecommendations();
    console.log("fecthed data", data);
    return(
        data
    );
}

export const BestForYou = () => {
    return(
        <div className= {styles.comp}>
            <div className= {styles.card}>
                <div className= {styles.textbox}>{test.distance}</div>
                <div className= {styles.textbox}>{test.name}</div>
                <div className= {styles.textbox}>{test.img}</div>
                <div className= {styles.textbox}>{test.shipping}</div>
                <div className= {styles.textbox}>{test.shop}</div>
            </div>
        </div>
    );
}

export default WatchList;