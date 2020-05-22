import React from "react"
import { Link } from "gatsby"

import Layout from "@app/components/layout"
import Image from "@app/components/Image"
import SEO from "@app/components/seo"

const IndexPage = () => (
  <Layout>
    <SEO title="Home" description="hoge" lang="ja" />
    <h1>Hi people</h1>
    <p>Welcome to your new Gatsby site.</p>
    <p>Now go build something great.</p>
    <Image filename="image01.png" />
    <Image filename="image02.png" />
    <Image filename="image03.png" />
    <Image filename="image04.jpg" />
    <Image filename="image05.jpg" />
    <Image filename="image06.jpeg" />
    <Link to="/page-2/">Go to page 2</Link>
  </Layout>
)

export default IndexPage
