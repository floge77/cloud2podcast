import React from "react";

function App() {
  fetch("/podcasts")
    .then((data) => data.json())
    .then(console.log);

  return (
    <div className="App">
      <header className="App-header">
        <h1>Cloud2Podcasts</h1>
      </header>
    </div>
  );
}

export default App;
