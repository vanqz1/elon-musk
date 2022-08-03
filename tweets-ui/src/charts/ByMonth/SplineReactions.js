import React from 'react';
import Highcharts from 'highcharts';
import HighchartsReact from 'highcharts-react-official';
import {useState, useEffect} from "react";

function SplineReactions(){
    const [tweets, setTweets] = useState([]);
    let xAxisLabel = tweets.map((tweet) => tweet.label);
    let yAxisDataRetweets = tweets.map((tweet) => Math.round(tweet.averageRetweets));
    let yAxisDataLikes = tweets.map((tweet) => Math.round(tweet.averageLikes));
    let yAxisDataRelies = tweets.map((tweet) => Math.round(tweet.averageReplies));

    const avgLikesPerMonth = xAxisLabel.map((child,index) => {
        const newDate =labelToDate(child)
        const month = newDate.getMonth();
        const year = newDate.getFullYear();
        return [Date.UTC(year, month),yAxisDataLikes[index]]
    });

    const avgRetweetsPerMonth = xAxisLabel.map((child,index) => {
        const newDate = labelToDate(child)
        const month = newDate.getMonth();
        const year = newDate.getFullYear();
        return [Date.UTC(year, month),yAxisDataRetweets[index]]
    });

    const avgRepliesPerMonth = xAxisLabel.map((child,index) => {
        const newDate = labelToDate(child)
        const month = newDate.getMonth();
        const year = newDate.getFullYear();
        return [Date.UTC(year, month),yAxisDataRelies[index]]
    });

    const options = {
        chart: {
            type: 'spline'
        },
        xAxis: {
            type: 'datetime',
            title: {
                text: 'month/year'
            },
            tickPixelInterval: 2
        },
        yAxis: {
            type: 'number',
            title: {
                text: 'Number'
            },
            tickPixelInterval: 10
        },
        title: {
            text: 'Average likes, retweets, replies per month over the years'
        },
        series: [
            {
                data: avgLikesPerMonth,
                name: "Average likes",
            },
            {
                data: avgRetweetsPerMonth,
                name: "Average retweets",
                color: '#32a838',
            },
            {
                data: avgRepliesPerMonth,
                name: "Average replies",
                color: '#a84632',
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

export default SplineReactions;