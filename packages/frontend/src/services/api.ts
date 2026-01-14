export const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api';

export interface Anime {
    id: string;
    title: string;
    image: string;
    last_episode?: string;
    last_episode_id?: string;
    status?: string;
    rating?: string;
    type?: string;
    genre?: string[];
}

export interface Episode {
    id: string;
    title: string;
    date: string;
}

export interface AnimeDetail extends Anime {
    synopsis: string;
    episodes: Episode[];
}

export interface StreamLink {
    quality: string;
    url: string;
}

export interface StreamResponse {
    title: string;
    link: string;
    qualities: StreamLink[];
}

export const apiService = {
    async getHome(): Promise<Anime[]> {
        const response = await fetch(`${API_BASE_URL}/home`);
        if (!response.ok) throw new Error('Failed to fetch home data');
        return response.json();
    },

    async getOngoing(): Promise<Anime[]> {
        const response = await fetch(`${API_BASE_URL}/ongoing`);
        if (!response.ok) throw new Error('Failed to fetch ongoing data');
        return response.json();
    },

    async getPopular(): Promise<Anime[]> {
        const response = await fetch(`${API_BASE_URL}/popular`);
        if (!response.ok) throw new Error('Failed to fetch popular data');
        return response.json();
    },

    async searchAnime(query: string): Promise<Anime[]> {
        const response = await fetch(`${API_BASE_URL}/search?q=${encodeURIComponent(query)}`);
        if (!response.ok) throw new Error('Failed to search anime');
        return response.json();
    },

    async getAnimeDetail(id: string): Promise<AnimeDetail> {
        const response = await fetch(`${API_BASE_URL}/anime/${id}`);
        if (!response.ok) throw new Error('Failed to fetch anime detail');
        return response.json();
    },

    async getEpisodeDetail(id: string): Promise<StreamResponse> {
        const response = await fetch(`${API_BASE_URL}/episode/${id}`);
        if (!response.ok) throw new Error('Failed to fetch episode detail');
        return response.json();
    },

    async getGenres(): Promise<string[]> {
        const response = await fetch(`${API_BASE_URL}/genres`);
        if (!response.ok) throw new Error('Failed to fetch genres');
        return response.json();
    },

    normalizeUrl(url: string): string {
        if (url.startsWith('/api/')) {
            // Remove /api/ from start as API_BASE_URL already includes it
            return `${API_BASE_URL}${url.substring(4)}`;
        }
        return url;
    }
};
