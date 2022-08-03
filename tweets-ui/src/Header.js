import React from 'react';

function Header(){
    return (
        <nav className="navbar navbar-expand-lg navbar-light bg-light">
            <a className="navbar-brand" href="/">TweetsCharts</a>
            <button className="navbar-toggler" type="button" data-toggle="collapse"
                    aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                <span className="navbar-toggler-icon"></span>
            </button>
            <div className="collapse navbar-collapse" id="navbarNav">
                <ul className="navbar-nav">
                    <li className="nav-item active">
                        <a className="nav-link" href="reactions-per-month">Average Reactions Per Month</a>
                    </li>
                    <li className="nav-item">
                        <a className="nav-link" href="tweets-per-month">Tweets Per Month</a>
                    </li>
                    <li className="nav-item">
                        <a className="nav-link" href="tweets-per-hour">Tweets Per Time Of Day</a>
                    </li>
                    <li className="nav-item">
                        <a className="nav-link" href="tweets-per-year">Tweets Per Year</a>
                    </li>
                </ul>
            </div>
        </nav>
    )
}

export default Header;