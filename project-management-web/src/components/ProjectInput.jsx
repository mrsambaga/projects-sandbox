import React, { useState } from "react";
import Button from "./Button";
import Form from "./Form";

export default function ProjectInput({ addProject }) {
  const [project, setProject] = useState({
    title: "",
    date: "",
    description1: "",
    description2: "",
  });

  function handleSubmit() {
    addProject(project);

    setProject({
      title: "",
      date: "",
      description1: "",
      description2: "",
    });
  }

  function handleFormChange(event) {
    const { name, value } = event.target;
    setProject({
      ...project,
      [name]: value,
    });
  }

  return (
    <div className="w-1/2 flex flex-col space-y-5 px-10 py-16">
      <h1 className="text-3xl font-bold text-stone-600 mb-5 text-center">
        Add New Project
      </h1>
      <Form title="Project Title" field={"title"} onChange={handleFormChange} />
      <Form title="Start Date" field={"date"} onChange={handleFormChange} />
      <Form
        title="Description 1"
        field={"description1"}
        onChange={handleFormChange}
      />
      <Form
        title="Description 2"
        field={"description2"}
        onChange={handleFormChange}
      />
      <div className="flex justify-end">
        <Button
          title="Submit"
          style="text-lg font-bold text-stone-600 w-1/5 border border-black py-2 px-4 rounded hover:bg-stone-100"
          clickHandler={handleSubmit}
        />
      </div>
    </div>
  );
}
