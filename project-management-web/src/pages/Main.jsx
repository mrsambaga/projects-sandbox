import React, { useState } from "react";
import SideBar from "../components/SideBar";
import Project from "../components/Project";
import ProjectInput from "../components/ProjectInput";

export default function Main() {
  const [isInputProject, setInputProject] = useState(true);
  const [projects, setProjects] = useState([]);

  function handleAddProject(newProject) {
    setProjects([...projects, newProject]);
  }

  return (
    <div className="flex h-full w-full">
      <SideBar />
      {isInputProject ? (
        <ProjectInput addProject={handleAddProject} />
      ) : (
        <Project />
      )}
    </div>
  );
}
