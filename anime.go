package anime

import (
	"fmt"

	"github.com/nokusukun/jikan2go/anime"
)

func main() {
	jikan.User{Username: "ivanovishado", Data: "animelist", AnimeListFilter: "completed"}

	anime, _ := jikan.Anime{ID: 1}.Get()

	// You can print values of type interface{}
	fmt.Println(anime["title"])
}

// import json
// from jikanpy import Jikan

// jikan = Jikan(selected_base="http://localhost:8080/v3/")

// data = jikan.user(username="ivanovishado", request="animelist", argument="completed")

// # Traer score, rating, start/end date, title, episodes, type, airing status, url, mal_id
// # Anime endpoint: stats

// anime = data["anime"]
// anime_list = []

// for a in anime:
//     mal_id = int(a["mal_id"])
//     anime_info = jikan.anime(mal_id)

//     anime_list.append(
//         [
//             mal_id,
//             a["title"],
//             a["score"],
//             a["rating"],
//             a["start_date"],
//             a["end_date"],
//             a["watch_start_date"],
//             a["watch_end_date"],
//             a["total_episodes"],
//             a["type"],
//             a["airing_status"],
//             a["url"],
//             anime_info["airing"],
//             anime_info["duration"],
//             anime_info["ending_themes"],
//             anime_info["episodes"],
//             anime_info["favorites"],
//             # anime_info["genres"]["name"],
//             anime_info["members"],
//             anime_info["opening_themes"],
//             anime_info["popularity"],
//             anime_info["rank"],
//             anime_info["score"],
//             anime_info["scored_by"],
//             anime_info["source"],
//             anime_info["status"],
//             # anime_info["studios"],
//             anime_info["synopsis"],
//         ]
//     )

// with open('animelist.json', 'w') as json_file:
//     json.dump(anime_list, json_file)
