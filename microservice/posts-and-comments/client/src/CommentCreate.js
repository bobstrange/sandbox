import React, { useState } from "react";
import axios from "axios";

export default ({ postId }) => {
  const [content, setContent] = useState("");

  const handleChange = (e) => {
    setContent(e.target.value);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    await axios.post(`http://comments-cluster-ip-srv:8081/posts/${postId}/comments`, { content });

    setContent('')
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label>New Comment</label>
          <input
            value={content}
            onChange={handleChange}
            className="form-control"
          />
        </div>
        <button className="btn btn-primary">Submit</button>
      </form>
    </div>
  );
};
