import React, { useState, useEffect } from "react";

export default function Search() {
const [movies, setMovies] = useState([]);
const [loading, setLoading] = useState(false);
const [query, setQuery] = useState("");

const fetchMovies = async () => {
  setLoading(true);
  try {
    const response = await fetch(`http://localhost:8080/search?moviename=${query}`);
    if (!response.ok) throw new Error("Failed to fetch movies");

    const data = await response.json();
    console.log("API Response:", data);

    setMovies(data.results ?? []);
  } catch (error) {
    console.error("Error fetching movies:", error);
    setMovies([]);
  } finally {
    setLoading(false);
  }
};

const addMovie = async (movie) => {
  try {
    const response = await fetch("http://localhost:8080/post", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(movie),
    });

    if (!response.ok) {
      throw new Error("Failed to add the movie");
    }

    const result = await response.json();
    console.log("Movie added:", result);

  } catch (error) {
    console.error("Error adding movie:", error);
  }
};

useEffect(() => {
  fetchMovies();
}, []);

return (
  <div>
    <h1>Movie Search</h1>
    <br />
    <input
      type="text"
      className="search"
      value={query}
      onChange={(e) => setQuery(e.target.value)}
      placeholder="Enter movie name"
    />
    <button onClick={fetchMovies}>Search</button>
    <br />
    <br />
    <a href="/"><button>Return home</button></a>

    {loading && <p>Loading...</p>}

    {movies.length === 0 && !loading && <p>No movies found</p>}

    <ul>
      {movies.map((movie) => (
        
        <div key={movie.id}>
        <br />
        --------------------------------------------------------------------------------
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
          <button onClick={() => addMovie(movie)}>ADD</button>
        </div>
      ))}
    </ul>
  </div>
);
}
