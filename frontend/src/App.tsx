import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import PageNotFound from "./pages/not-found";
import Layout from "./components/layout";
import AllPost from "./pages/all-post";
import PreviewPost from "./pages/preview-post";

function App() {
  return (
    <>
      <Router>
        <Routes>
          <Route path="*" element={<PageNotFound />} />
          <Route path="/" element={<Layout />}>
            <Route index element={<AllPost />} />
            <Route path="/preview" element={<PreviewPost />} />
          </Route>
        </Routes>
      </Router>
    </>
  );
}

export default App;
