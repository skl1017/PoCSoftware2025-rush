export default function Home() {
    return (
    <div className="Home">
        <h1>My Watchlist</h1>
        <p>Search and add your favorite film to your watch list</p>
        <br /><br />
        <a href="/Search"><button>
          Search
        </button></a>&emsp;|&emsp;
        <a href="/Watchlist"><button>
          Watchlist
        </button></a>
   </div> 
   );
}
