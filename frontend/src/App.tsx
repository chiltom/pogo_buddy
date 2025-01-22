import { ReactElement } from "react";
import { Outlet } from "react-router-dom";
import "./App.css";

const App: React.FC = (): ReactElement => {
  return (
    <div className="App">
      <header className="App-header">
        <Outlet />
      </header>
    </div>
  );
};

export default App;
