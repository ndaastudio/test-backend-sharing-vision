import React, { useEffect, useState } from "react";
import axios from "axios";
import { Button } from "@/components/ui/button";

type Article = {
  id: number;
  title: string;
  content: string;
  category: string;
  status: "Publish" | "Draft" | "Trash";
};

const PreviewPost: React.FC = () => {
  const [articles, setArticles] = useState<Article[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  const [limit] = useState<number>(5);
  const [offset] = useState<number>(0);
  const [page, setPage] = useState<number>(1);

  const fetchArticles = async () => {
    try {
      setLoading(true);
      const res = await axios.get(
        `http://localhost:4000/article/${limit}/${offset}?status=Publish`
      );

      setArticles(res.data.data);
    } catch (err) {
      setError("Gagal memuat artikel");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchArticles();
  }, [offset, limit]);

  return (
    <div className="max-w-3xl mx-auto p-6">
      <h1 className="text-3xl font-bold mb-6">Blog</h1>

      {loading && <p>Loading...</p>}
      {error && <p className="text-red-500">{error}</p>}

      <div className="space-y-6">
        {articles.map((article) => (
          <div
            key={article.id}
            className="p-4 border rounded-lg shadow-sm hover:shadow-md transition"
          >
            <h2 className="text-xl font-semibold">{article.title}</h2>
            <p className="text-sm text-gray-500">{article.category}</p>
            <p className="mt-2 text-gray-700">
              {article.content.slice(0, 120)}...
            </p>
          </div>
        ))}
      </div>

      <div className="flex items-center justify-between">
        <Button
          variant="outline"
          disabled={page === 1}
          onClick={() => setPage((prev) => Math.max(prev - 1, 1))}
        >
          Previous
        </Button>
        <span>Page {page}</span>
        <Button variant="outline" onClick={() => setPage((prev) => prev + 1)}>
          Next
        </Button>
      </div>
    </div>
  );
};

export default PreviewPost;
