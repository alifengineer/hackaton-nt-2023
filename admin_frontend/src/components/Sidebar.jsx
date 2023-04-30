import React from "react";

const Sidebar = ({ handleChange, menus, active }) => {
  return (
    <div style={{ height: "100%" }} className="sidebar">
      <nav class="navbar navbar-expand-lg bg-body-tertiary">
        <div class="container-fluid">
          <div class="navbar-nav d_flex flex flex-column mb-3">
            {menus?.map((menu, index) => {
              return (
                <a
                  key={index}
                  class={`nav-link ${menu === active ? "active" : ""}}`}
                  href="#"
                  onClick={() => handleChange(menu)}
                >
                  {menu}
                </a>
              );
            })}
          </div>
        </div>
      </nav>
    </div>
  );
};

export default Sidebar;
