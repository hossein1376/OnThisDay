/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  images: {
    remotePatterns: [
      {
        protocol: "https",
        hostname: "upload.wikimedia.org/",
        port: "",
        pathname: "/wikipedia/commons/thumb/*",
      },
    ],
  },
};

module.exports = nextConfig;
