import React from "react";
import SideBar from "../components/SideBar";
import Project from "../components/Project";

export default function Main() {
  return (
    <div className="flex h-full w-full">
      <SideBar />
      <Project />
    </div>
  );
}
