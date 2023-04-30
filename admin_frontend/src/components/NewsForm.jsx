import React, { useState } from "react";
import axios from "../utils/axios";
import { simpleAxios } from "../utils/axios";
import { toast } from "react-hot-toast";

const NewsForm = () => {
  const [selectedImage, setSelectedImage] = useState(null);
  const [title, setTitle] = useState("");
  const [text, setText] = useState("");

  function handleImageChange(event) {
    const imageFile = event.target.files[0];
    setSelectedImage(imageFile);
  }

  const handleFormSubmit = async (event) => {
    event.preventDefault();
    const formData = new FormData();
    formData.append("file", selectedImage);
    const imgUrl = await simpleAxios
      .post("https://dilmurodov.jprq.live/api/v1/image/upload", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      })
      .then((response) => {
        return response.data.data.url;
      })
      .catch((error) => {
        console.log(error);
      });
    await axios
      .post(`/api/v1/news/`, {
        title: title,
        content: text,
        image: imgUrl,
      })
      .then((res) => {
        toast.success("News Created");
        setText("");
        setTitle("");
        setSelectedImage(null);
        console.log(res);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  return (
    <div
      className="news_form"
      style={{ marginLeft: "30px", marginTop: "50px" }}
    >
      <form onSubmit={handleFormSubmit} className="needs-validation" novalidate>
        <div class="mb-3">
          <label htmlFor="exampleFormControlTextarea1" class="form-label">
            News title
          </label>
          <input
            class="form-control"
            type="text"
            placeholder="Title"
            aria-label="title"
            required={true}
            onChange={(e) => setTitle(e.target.value)}
          />
        </div>
        <div class="mb-3">
          <label htmlFor="exampleFormControlTextarea1" class="form-label">
            Example textarea
          </label>
          <textarea
            onChange={(e) => setText(e.target.value)}
            class="form-control"
            id="exampleFormControlTextarea1"
            rows="3"
          ></textarea>
        </div>
        <div class="mb-3">
          <label htmlFor="formFile" class="form-label">
            Image
          </label>
          <input
            class="form-control"
            type="file"
            accept="image/*"
            id="formFile"
            onChange={handleImageChange}
            required={true}
          />
        </div>
        <button type="submit" class="btn btn-success">
          Save
        </button>
      </form>
    </div>
  );
};

export default NewsForm;
