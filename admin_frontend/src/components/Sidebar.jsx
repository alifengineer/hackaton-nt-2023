import React from "react";

const Sidebar = () => {
  return (
    <div style={{ height: "100%" }} className="sidebar">
      <nav class="navbar navbar-expand-lg bg-body-tertiary">
        <div class="container-fluid">
          <div class="navbar-nav d_flex flex flex-column mb-3">
            <a class="nav-link active" aria-current="page" href="#">
              Match
            </a>
            <a class="nav-link" href="#">
              News
            </a>
          </div>
        </div>
      </nav>
    </div>
  );
};

export default Sidebar;
