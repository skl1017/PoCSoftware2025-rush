import Home from '../pages/Home'
import Search from '../pages/search'
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";


export default function App() {
    return (
      <>
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/search" element={<Search />} />
        </Routes>
      </Router>
      </>
    );
  }