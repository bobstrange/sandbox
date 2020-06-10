import React from 'react'
import PostCreate from './PostCreate'
import PostList from './PostList'

export default () => {
  return <div className="container">
    <h1>Posts</h1>
    <h2>Create Post</h2>
    <PostCreate />
    <h2>Posts</h2>
    <PostList />
  </div>
}

