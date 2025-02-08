import React, { useState, useEffect } from "react";

export default function Search() {
const [movies, setMovies] = useState([]);


const fetchMovies = async () => {
try {
    const response = await fetch(`http://localhost:8080/watchlist`);
    if (!response.ok) throw new Error("Failed to fetch movies");

    const data = await response.json();
    console.log("API Response:", data);

    setMovies(data.results ?? []);
} catch (error) {
    console.error("Error fetching movies:", error);
    setMovies([]);
}
};

useEffect(() => {
fetchMovies();
}, []);

return (
<div>
    <h1>My Watchlist</h1>
    <ul>
    {movies.map((movie) => (
        <div key={movie.id}>
        <h2>{movie.title}</h2>
        
        {movie.poster_path && (
          <img
            src={`https://image.tmdb.org/t/p/w500${movie.poster_path}`} 
            alt={movie.title}
            style={{ width: "200px", height: "auto" }}
          />
        )}
        <p>{movie.overview}</p>
        <br />
        <button>DELL</button>
        <br />
        <br />
        --------------------------------------------------------------------------------
      </div>
    ))}
    </ul>
    <a href="/"><button>Return Home</button></a>
</div>
);
}