import { useState } from "react";
import axios from "axios";

interface DeleteModalProps {
  isOpen: boolean;
  onClose: () => void;
  articleId: number | null;
  onDeleted: () => void;
}

export default function DeleteArticleModal({
  isOpen,
  onClose,
  articleId,
  onDeleted,
}: DeleteModalProps) {
  const [loading, setLoading] = useState(false);

  const handleDelete = async () => {
    if (!articleId) return;
    try {
      setLoading(true);
      await axios.put(`http://localhost:4000/article/${articleId}`, {
        status: "Trash",
      });
      onDeleted();
      onClose();
    } catch (err) {
      console.error("Gagal menghapus artikel", err);
    } finally {
      setLoading(false);
    }
  };

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-black/40 z-50">
      <div className="bg-white rounded-xl shadow-lg p-6 w-full max-w-md">
        <h2 className="text-xl font-semibold text-gray-800 mb-4">
          Delete Article
        </h2>
        <p className="text-gray-600 mb-6">
          Are you sure you want to move this article to <b>Trash</b>?
        </p>

        <div className="flex justify-end gap-3">
          <button
            onClick={onClose}
            disabled={loading}
            className="px-4 py-2 rounded-lg border border-gray-300 text-gray-700 hover:bg-gray-100"
          >
            Cancel
          </button>
          <button
            onClick={handleDelete}
            disabled={loading}
            className="px-4 py-2 rounded-lg bg-red-600 text-white hover:bg-red-700 disabled:opacity-50"
          >
            {loading ? "Deleting..." : "Move to Trash"}
          </button>
        </div>
      </div>
    </div>
  );
}
