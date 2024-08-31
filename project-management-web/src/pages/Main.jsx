import React, { useState } from "react";
import SideBar from "../components/SideBar";
import Project from "../components/Project";
import ProjectInput from "../components/ProjectInput";

export default function Main() {
  const [isInputProject, setInputProject] = useState(true);
  const [projects, setProjects] = useState([]);
  const [selectedProject, setSelectedProject] = useState({
    title: "",
    date: "",
    description1: "",
    description2: "",
  });

  function handleAddProject(newProject) {
    setProjects([...projects, newProject]);
  }

  function handleSidebarClick(title) {
    if (title == "+ Add Project") {
      setInputProject(true);
    } else {
      const activeProject = projects.find((project) => project.title === title);
      setSelectedProject(activeProject);
      setInputProject(false);
    }
  }

  return (
    <div className="flex h-full w-full">
      <SideBar projects={projects} onClickButton={handleSidebarClick} />
      {isInputProject ? (
        <ProjectInput addProject={handleAddProject} />
      ) : (
        <Project activeProject={selectedProject} />
      )}
    </div>
  );
}
