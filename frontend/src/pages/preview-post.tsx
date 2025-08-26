import { useState } from "react";
import axios from "axios";
import { Button } from "@/components/ui/button";
import { useQuery } from "@tanstack/react-query";

type Post = {
  id: number;
  title: string;
  content: string;
  category: string;
  status: "Draft" | "Publish";
  created_date: string;
  updated_date: string;
};

type ApiResponse = {
  data: Post[];
  success: boolean;
};

async function fetchPosts(limit: number, offset: number) {
  const res = await axios.get<ApiResponse>(
    `http://localhost:4000/article/${limit}/${offset}?status=Publish`
  );
  return res.data.data;
}

export default function PreviewPost() {
  const [page, setPage] = useState(1);
  const limit = 3;
  const offset = (page - 1) * limit;

  const { data, isLoading, error, isError } = useQuery<Post[]>({
    queryKey: ["posts", limit, offset],
    queryFn: () => fetchPosts(limit, offset),
    placeholderData: (previousData) => previousData,
  });

  if (isLoading) {
    return <p className="p-4 text-center">Loading...</p>;
  }

  if (isError) {
    return (
      <p className="p-4 text-red-500">
        Gagal mengambil data:{" "}
        {error instanceof Error ? error.message : "Error tidak diketahui"}
      </p>
    );
  }

  return (
    <div className="max-w-3xl mx-auto p-6">
      <h1 className="text-3xl font-bold mb-6">Posts</h1>

      {isLoading && <p>Loading...</p>}
      {isError && <p className="text-red-500">{error}</p>}

      <div className="space-y-6">
        {data?.map((post) => (
          <div
            key={post.id}
            className="p-4 border rounded-lg shadow-sm hover:shadow-md transition"
          >
            <h2 className="text-xl font-semibold">{post.title}</h2>
            <p className="text-sm text-gray-500">{post.category}</p>
            <p className="mt-2 text-gray-700">
              {post.content.slice(0, 120)}...
            </p>
          </div>
        ))}
      </div>

      <div className="flex items-center justify-between mt-4">
        <Button
          variant="outline"
          disabled={page === 1}
          onClick={() => setPage((prev) => Math.max(prev - 1, 1))}
        >
          Previous
        </Button>
        <span>Page {page}</span>
        <Button
          disabled={(data?.length ?? 0) < limit}
          variant="outline"
          onClick={() => setPage((prev) => prev + 1)}
        >
          Next
        </Button>
      </div>
    </div>
  );
}
