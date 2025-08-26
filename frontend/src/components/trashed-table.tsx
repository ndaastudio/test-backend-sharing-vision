import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Button } from "./ui/button";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";
import { useState } from "react";
import EditArticleModal from "./edit-article-modal";

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
    `http://localhost:4000/article/${limit}/${offset}?status=Trash`
  );
  return res.data.data;
}

export default function TrashedTable() {
  const [page, setPage] = useState(1);
  const limit = 5;
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
    <div className="space-y-4">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[100px]">No.</TableHead>
            <TableHead>Title</TableHead>
            <TableHead>Category</TableHead>
            <TableHead>Action</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {data?.map((post, index) => (
            <TableRow key={post.id}>
              <TableCell className="font-medium">
                {offset + index + 1}
              </TableCell>
              <TableCell>{post.title}</TableCell>
              <TableCell>{post.category}</TableCell>
              <TableCell>
                <div className="flex justify-start gap-4">
                  <EditArticleModal article={post} />
                </div>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>

      <div className="flex items-center justify-between">
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
