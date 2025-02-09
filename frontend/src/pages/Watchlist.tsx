import React, { useState, useEffect } from "react";

export default function Watchlist() {
const [movies, setMovies] = useState([]);

const transformData = (apiData) => {
return apiData.map((movieArray) => {
    const movieObject = {};
    movieArray.forEach(({ Key, Value }) => {
    movieObject[Key] = Value;
    });
    return movieObject;
});
};

const fetchMovies = async () => {
try {
    const response = await fetch(`http://localhost:8080/watchlist`);
    if (!response.ok) throw new Error("Failed to fetch movies");

    const data = await response.json();
    console.log("API Response:", data);

    setMovies(transformData(data));
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
        <br />
        <a href="/">
        <button>Return Home</button>
        </a>
        
        {movies.length === 0 && <p>No movies in your watchlist.</p>}
        <ul>
        <br />
        <br />
            --------------------------------------------------------------------------------
        {movies.map((movie) => (
            <div key={movie._id}>
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
            <button>DEL</button>
            </div>
        ))}
        </ul>

    </div>
);
}
