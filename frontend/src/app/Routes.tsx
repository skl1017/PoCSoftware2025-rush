import Home from '../pages/Home'
import Search from '../pages/search'
import Watchlist from '../pages/Watchlist'
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";


export default function App() {
    return (
      <>
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/search" element={<Search />} />
          <Route path="/watchlist" element={<Watchlist />} />
        </Routes>
      </Router>
      </>
    );
  }