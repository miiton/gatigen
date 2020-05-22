module.exports = {
  siteMetadata: {
    title: `ガチジェン`,
    description: `A code-generator for gatsby-image`,
    author: `@miiton`,
    siteUrl: `https://github.com/miiton/gatigen`,
  },
  plugins: [
    `gatsby-plugin-typescript`,
    {
      resolve: `gatsby-source-filesystem`,
      options: {
        name: `images`,
        path: `${__dirname}/src/images`,
      },
    },
    `gatsby-transformer-sharp`,
    `gatsby-plugin-sharp`,
  ],
}
