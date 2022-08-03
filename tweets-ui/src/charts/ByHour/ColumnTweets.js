import React from 'react';
import Highcharts from 'highcharts';
import HighchartsReact from 'highcharts-react-official';
import {useState, useEffect} from "react";

function ColumnTweets(){
    const [tweets, setTweets] = useState([]);
    let xAxisLabel = tweets.map((tweet) => tweet.label);
    let yAxisDataTweetsCount = tweets.map((tweet) => Math.round(tweet.tweetsCount));

    const tweetsCountPerMonth = xAxisLabel.map((child,index) => {
        return [getLabelText(child),yAxisDataTweetsCount[index]]
    });

    const options = {
        chart: {
            type: 'column'
        },
        xAxis: {
            type: 'string',
            title: {
                text: 'o\'clock'
            },
            tickPixelInterval: 1
        },
        yAxis: {
            type: 'number',
            title: {
                text: 'Number'
            },
            tickPixelInterval: 100
        },
        title: {
            text: 'Tweets count per time of day'
        },
        series: [
            {
                data: tweetsCountPerMonth,
                name: "Tweets count",
            }
        ]
    };

    useEffect(function (){
        fetch('http://localhost:8030/api/tweets/aggregate?groupBy=hour')
            .then((response) => response.json())
            .then((json) => setTweets(json));
    },[])

    return (
        <div>
            <HighchartsReact highcharts={Highcharts} options={options} />
        </div>
    )
}

function getLabelText(label){
    const result = label.replace( "-hour", " o'clock")

    return result
}

export default ColumnTweets;