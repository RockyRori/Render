import axios from "axios";

const API_BASE_URL = "http://localhost:8080";

export interface Game {
  id: number;
  title: string;
  price: number;
  release_date?: string; // 可能后端没有返回这个字段
}

export async function fetchGames(): Promise<Game[]> {
  try {
    const response = await axios.get(`${API_BASE_URL}/games`);
    return response.data.map((game: any) => ({
      id: game.id,
      title: game.title,
      price: game.price,
      description: game.description,
      release_date: game.release_date || "未知"
    }));
  } catch (error) {
    console.error("获取游戏数据失败:", error);
    return [];
  }
}
