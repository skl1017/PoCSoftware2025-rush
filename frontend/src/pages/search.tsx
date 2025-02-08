export default function Search() {
    return (
      <div className="SearchBox">
        <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"></link>
        <label>
            <input type="text" className="search" defaultValue="Search"/>
        </label>
        <button type="submit"><i className="fa fa-search"></i></button>
      </div>
    );
  }
  