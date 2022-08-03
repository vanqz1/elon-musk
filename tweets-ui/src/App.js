import './App.css';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import Header from "./Header";
import SplineReactions from "./charts/ByMonth/SplineReactions";
import ColumnTweetsPerMonth from "./charts/ByMonth/ColumnTweets";
import ColumnTweetsPerHour from "./charts/ByHour/ColumnTweets";
import ColumnTweetsPerYear from "./charts/ByYear/ColumnTweets";

function App() {
  return (
      <div className="container-fluid">
          <div className="row">
                <div className="App">
                  <Header />
                    <Router>
                        <Routes>
                            <Route path="/" element={<SplineReactions/>}></Route>
                            <Route path="/reactions-per-month" element={<SplineReactions/>}></Route>
                            <Route path="/tweets-per-month" element={<ColumnTweetsPerMonth/>}></Route>
                            <Route path="/tweets-per-hour" element={<ColumnTweetsPerHour/>}></Route>
                            <Route path="/tweets-per-year" element={<ColumnTweetsPerYear/>}></Route>
                        </Routes>
                    </Router>
                </div>
          </div>
      </div>
  );
}

export default App;
