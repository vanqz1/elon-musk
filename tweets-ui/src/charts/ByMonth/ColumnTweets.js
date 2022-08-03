import React from 'react';
import Highcharts from 'highcharts';
import HighchartsReact from 'highcharts-react-official';
import {useState, useEffect} from "react";

function ColumnTweets(){
    const [tweets, setTweets] = useState([]);
    let xAxisLabel = tweets.map((tweet) => tweet.label);
    let yAxisDataTweetsCount = tweets.map((tweet) => Math.round(tweet.tweetsCount));

    const tweetsCountPerMonth = xAxisLabel.map((child,index) => {
        const newDate = new labelToDate(child)
        const month = newDate.getMonth();
        const year = newDate.getFullYear();
        return [Date.UTC(year, month),yAxisDataTweetsCount[index]]
    });

    const options = {
        chart: {
            type: 'column'
        },
        xAxis: {
            type: 'datetime',
            title: {
                text: 'month/year'
            },
            tickPixelInterval: 1
        },
        yAxis: {
            type: 'number',
            title: {
                text: 'Number'
            },
            tickPixelInterval: 30
        },
        title: {
            text: 'Tweets count per month over the years'
        },
        series: [
            {
                data: tweetsCountPerMonth,
                name: "Tweets count",
            }
        ]
    };

    useEffect(function (){
        fetch('http://localhost:8030/api/tweets/aggregate?groupBy=year&groupBy=month')
            .then((response) => response.json())
            .then((json) => setTweets(json));
    },[])

    return (
        <div>
            <HighchartsReact highcharts={Highcharts} options={options} />
        </div>
    )
}

function labelToDate(label){
    const cleanString = label.replace( "-year", "").replace("-month", "");
    const modString = cleanString.replace(/ /, " 1 ");
    const date = new Date(modString);

    return date
}

export default ColumnTweets;