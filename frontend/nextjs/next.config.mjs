/** @type {import('next').NextConfig} */
const nextConfig = {
    async rewrites() {
        return [
               {
                     source: '/api/v1/users/:path*',
                     destination: `http://localhost:10002/api/v1/users/:path*`,
               },
        ]
    },
};

export default nextConfig;
