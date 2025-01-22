import { createBrowserRouter } from "react-router-dom";
import App from "./App";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      {
        index: true,
        element: <div>Hello World</div>,
      },
    ],
    errorElement: <div>404 Not Found</div>,
  },
]);

export default router;
