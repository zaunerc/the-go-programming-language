package model

import "time"

type SearchResponse struct {
	Autofilter struct {
		Enable    bool   `json:"enable"`
		Apply     bool   `json:"apply"`
		Fieldlist string `json:"fieldlist"`
		Fields    struct {
			All struct {
				Q     string        `json:"q"`
				Fq    []interface{} `json:"fq"`
				Apply bool          `json:"apply"`
			} `json:"all"`
		} `json:"fields"`
		Matches struct {
			Blumen struct {
				Fields struct {
					SemCategoryNamesAll struct {
						Fq []string `json:"fq"`
					} `json:"sem_category_names_all"`
				} `json:"fields"`
				SemActions struct {
					Exclude struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"exclude"`
				} `json:"sem_actions"`
			} `json:"Blumen"`
		} `json:"matches"`
	} `json:"autofilter"`
	Response struct {
		NumFound int `json:"numFound"`
		Start    int `json:"start"`
		Docs     []struct {
			ID                     string      `json:"id"`
			Type                   string      `json:"type"`
			AppKey                 []string    `json:"appKey"`
			SiteHash               []string    `json:"siteHash"`
			UID                    []string    `json:"uid"`
			Pid                    []string    `json:"pid"`
			VariantID              []string    `json:"variantId"`
			Created                []time.Time `json:"created"`
			Changed                []time.Time `json:"changed"`
			Access                 []string    `json:"access"`
			SemCreationDate        time.Time   `json:"sem_creation_date"`
			SemModificationDate    time.Time   `json:"sem_modification_date"`
			SemRecordType          string      `json:"sem_record_type"`
			SemRecordSubtype       string      `json:"sem_record_subtype"`
			SemGrouping            string      `json:"sem_grouping"`
			SemParentRecord        string      `json:"sem_parent_record"`
			SemTitle               string      `json:"sem_title"`
			TitleInitial           string      `json:"titleInitial"`
			SemTextSearch          []string    `json:"sem_text_search"`
			SemTextSuggest         []string    `json:"sem_text_suggest"`
			SemTitleAlpha          string      `json:"sem_title_alpha"`
			SemActive              bool        `json:"sem_active"`
			SemContent             []string    `json:"sem_content"`
			SemImageThumbnailURL   string      `json:"sem_image_thumbnail_url,omitempty"`
			SemImageURL            []string    `json:"sem_image_url"`
			SemAddressStreet       string      `json:"sem_address_street"`
			SemAddressPostalCode   string      `json:"sem_address_postal_code"`
			SemAddressCity         string      `json:"sem_address_city"`
			SemAddressCountry      string      `json:"sem_address_country"`
			SemLocation            string      `json:"sem_location"`
			SemLocationLatlng      string      `json:"sem_location_latlng"`
			SemContactWebsite      string      `json:"sem_contact_website"`
			SemContactEmail        []string    `json:"sem_contact_email,omitempty"`
			SemContactPhone        []string    `json:"sem_contact_phone"`
			SemSourceID            string      `json:"sem_source_id"`
			SemPopularity          int         `json:"sem_popularity"`
			SemContactName         string      `json:"sem_contact_name"`
			SemContactInfo         string      `json:"sem_contact_info"`
			SemURL                 string      `json:"sem_url"`
			Advantagecard          bool        `json:"advantagecard"`
			Bargainpriority        int         `json:"bargainpriority"`
			Companyuid             string      `json:"companyuid"`
			Customer               bool        `json:"customer"`
			Images                 []string    `json:"images"`
			Name                   string      `json:"name"`
			Section                []string    `json:"section"`
			Starttime              []time.Time `json:"starttime"`
			Endtime                []time.Time `json:"endtime"`
			Eventend               []time.Time `json:"eventend"`
			Eventstart             []time.Time `json:"eventstart"`
			EventstartFirst        time.Time   `json:"eventstart_first"`
			EventendLast           time.Time   `json:"eventend_last"`
			SemCategoryNamesBranch []string    `json:"sem_category_names_branch"`
			SemCategoryNamesAll    []string    `json:"sem_category_names_all"`
			SemCategoryIdsBranch   []string    `json:"sem_category_ids_branch"`
			SemTimestamp           time.Time   `json:"sem_timestamp"`
			Open                   bool        `json:"open"`
			SemExcerpt             string      `json:"sem_excerpt"`
			SemActions             struct {
				Default string `json:"default"`
				Goto    struct {
					SemURL        string `json:"sem_url"`
					SemURLAction  string `json:"sem_url_action"`
					SemURLLogging string `json:"sem_url_logging"`
					SemIsExternal bool   `json:"sem_is_external"`
				} `json:"goto"`
				Mlt struct {
					SemURL       string `json:"sem_url"`
					SemURLAction string `json:"sem_url_action"`
				} `json:"mlt"`
			} `json:"sem_actions"`
			SemEncodedID     string   `json:"sem_encoded_id"`
			SemContactMobile []string `json:"sem_contact_mobile,omitempty"`
			Dutyhours        []string `json:"dutyhours,omitempty"`
			SemOpeningHours  []string `json:"sem_opening_hours,omitempty"`
		} `json:"docs"`
	} `json:"response"`
	FacetCounts struct {
		FacetQueries struct {
			LocationFilters []interface{} `json:"location_filters"`
			TimeFilters     struct {
				Open struct {
					Value     string `json:"value"`
					Count     int    `json:"count"`
					SemAction struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"sem_action"`
				} `json:"open"`
				Today struct {
					Value     string `json:"value"`
					Count     int    `json:"count"`
					SemAction struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"sem_action"`
				} `json:"today"`
				Midday struct {
					Value     string `json:"value"`
					Count     int    `json:"count"`
					SemAction struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"sem_action"`
				} `json:"midday"`
				Evening struct {
					Value     string `json:"value"`
					Count     int    `json:"count"`
					SemAction struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"sem_action"`
				} `json:"evening"`
				Tomorrow struct {
					Value     string `json:"value"`
					Count     int    `json:"count"`
					SemAction struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"sem_action"`
				} `json:"tomorrow"`
				AfterTomorrow struct {
					Value     string `json:"value"`
					Count     int    `json:"count"`
					SemAction struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"sem_action"`
				} `json:"after_tomorrow"`
				Weekend struct {
					Value     string `json:"value"`
					Count     int    `json:"count"`
					SemAction struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"sem_action"`
				} `json:"weekend"`
			} `json:"time_filters"`
		} `json:"facet_queries"`
		FacetFields struct {
			AToZ struct {
				B struct {
					Value     string `json:"value"`
					Count     int    `json:"count"`
					SemAction struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"sem_action"`
				} `json:"b"`
				F struct {
					Value     string `json:"value"`
					Count     int    `json:"count"`
					SemAction struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"sem_action"`
				} `json:"f"`
				S struct {
					Value     string `json:"value"`
					Count     int    `json:"count"`
					SemAction struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"sem_action"`
				} `json:"s"`
			} `json:"a_to_z"`
			Section struct {
				Num1 struct {
					Value     string `json:"value"`
					Count     int    `json:"count"`
					SemAction struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"sem_action"`
				} `json:"1"`
				Num2 struct {
					Value     string `json:"value"`
					Count     int    `json:"count"`
					SemAction struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"sem_action"`
				} `json:"2"`
			} `json:"section"`
			Type struct {
				Tile struct {
					Value     string `json:"value"`
					Count     int    `json:"count"`
					SemAction struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"sem_action"`
				} `json:"tile"`
			} `json:"type"`
			Tiletype struct {
				BE struct {
					Value     string `json:"value"`
					Count     int    `json:"count"`
					SemAction struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"sem_action"`
				} `json:"BE"`
				SE struct {
					Value     string `json:"value"`
					Count     int    `json:"count"`
					SemAction struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"sem_action"`
				} `json:"SE"`
			} `json:"tiletype"`
			Tags []interface{} `json:"tags"`
		} `json:"facet_fields"`
		FacetDates struct {
		} `json:"facet_dates"`
		FacetRanges struct {
			Date struct {
				Counts        []interface{} `json:"counts"`
				GroupedCounts []interface{} `json:"groupedCounts"`
				Gap           string        `json:"gap"`
				Start         time.Time     `json:"start"`
				End           time.Time     `json:"end"`
				After         int           `json:"after"`
				Before        int           `json:"before"`
				Between       int           `json:"between"`
				SemAction     struct {
					SemURL       string `json:"sem_url"`
					SemURLAction string `json:"sem_url_action"`
				} `json:"sem_action"`
			} `json:"date"`
		} `json:"facet_ranges"`
		FacetIntervals struct {
		} `json:"facet_intervals"`
		FacetHeatmaps struct {
		} `json:"facet_heatmaps"`
		FacetPivot struct {
		} `json:"facet_pivot"`
		FacetHierarchies struct {
			Categories struct {
				Dienstleistungen struct {
					Value     string `json:"value"`
					Count     int    `json:"count"`
					SemAction struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"sem_action"`
					Options struct {
						Heiraten struct {
							Value     string `json:"value"`
							Count     int    `json:"count"`
							SemAction struct {
								SemURL       string `json:"sem_url"`
								SemURLAction string `json:"sem_url_action"`
							} `json:"sem_action"`
						} `json:"Heiraten"`
						Catering struct {
							Value     string `json:"value"`
							Count     int    `json:"count"`
							SemAction struct {
								SemURL       string `json:"sem_url"`
								SemURLAction string `json:"sem_url_action"`
							} `json:"sem_action"`
						} `json:"Catering"`
					} `json:"options"`
				} `json:"Dienstleistungen"`
				Shops struct {
					Value     string `json:"value"`
					Count     int    `json:"count"`
					SemAction struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"sem_action"`
					Options struct {
						WohnenHaushalt struct {
							Value     string `json:"value"`
							Count     int    `json:"count"`
							SemAction struct {
								SemURL       string `json:"sem_url"`
								SemURLAction string `json:"sem_url_action"`
							} `json:"sem_action"`
							Options struct {
								Garten struct {
									Value     string `json:"value"`
									Count     int    `json:"count"`
									SemAction struct {
										SemURL       string `json:"sem_url"`
										SemURLAction string `json:"sem_url_action"`
									} `json:"sem_action"`
									Options struct {
										Florist struct {
											Value     string `json:"value"`
											Count     int    `json:"count"`
											SemAction struct {
												SemURL       string `json:"sem_url"`
												SemURLAction string `json:"sem_url_action"`
											} `json:"sem_action"`
										} `json:"Florist"`
									} `json:"options"`
								} `json:"Garten"`
							} `json:"options"`
						} `json:"Wohnen & Haushalt"`
					} `json:"options"`
				} `json:"Shops"`
			} `json:"categories"`
			Region struct {
				Sterreich struct {
					Value     string `json:"value"`
					Count     int    `json:"count"`
					SemAction struct {
						SemURL       string `json:"sem_url"`
						SemURLAction string `json:"sem_url_action"`
					} `json:"sem_action"`
					Options struct {
						Enns struct {
							Value     string `json:"value"`
							Count     int    `json:"count"`
							SemAction struct {
								SemURL       string `json:"sem_url"`
								SemURLAction string `json:"sem_url_action"`
							} `json:"sem_action"`
						} `json:"Enns"`
					} `json:"options"`
				} `json:"Österreich"`
			} `json:"region"`
		} `json:"facet_hierarchies"`
	} `json:"facet_counts"`
	Spelling struct {
		Suggestions []struct {
			Suggestion string `json:"suggestion"`
			Hits       string `json:"hits"`
			SemAction  struct {
				SemURL       string `json:"sem_url"`
				SemURLAction string `json:"sem_url_action"`
			} `json:"sem_action"`
		} `json:"suggestions"`
	} `json:"spelling"`
	Sorting struct {
		SemAction struct {
			SemURL       string `json:"sem_url"`
			SemURLAction string `json:"sem_url_action"`
		} `json:"sem_action"`
		Fields struct {
			SemTitleAlpha struct {
				Asc struct {
					SemURL       string `json:"sem_url"`
					SemURLAction string `json:"sem_url_action"`
				} `json:"asc"`
				Desc struct {
					SemURL       string `json:"sem_url"`
					SemURLAction string `json:"sem_url_action"`
				} `json:"desc"`
			} `json:"sem_title_alpha"`
			Geodist struct {
				Asc struct {
					SemURL       string `json:"sem_url"`
					SemURLAction string `json:"sem_url_action"`
				} `json:"asc"`
				Desc struct {
					SemURL       string `json:"sem_url"`
					SemURLAction string `json:"sem_url_action"`
				} `json:"desc"`
			} `json:"geodist()"`
		} `json:"fields"`
	} `json:"sorting"`
	ActiveFilters struct {
		SemAction struct {
			SemURL       string `json:"sem_url"`
			SemURLAction string `json:"sem_url_action"`
		} `json:"sem_action"`
	} `json:"active_filters"`
	BaseUrls struct {
		SearchWithCurrentState struct {
			SemURL       string `json:"sem_url"`
			SemURLAction string `json:"sem_url_action"`
		} `json:"search_with_current_state"`
		SearchWithEmptyState struct {
			SemURL       string `json:"sem_url"`
			SemURLAction string `json:"sem_url_action"`
		} `json:"search_with_empty_state"`
		RemoveAllFilters struct {
			SemURL       string `json:"sem_url"`
			SemURLAction string `json:"sem_url_action"`
		} `json:"remove_all_filters"`
		RemoveSourceFilters struct {
			SemURL       string `json:"sem_url"`
			SemURLAction string `json:"sem_url_action"`
		} `json:"remove_source_filters"`
		DisableAutofilter struct {
			SemURL       string `json:"sem_url"`
			SemURLAction string `json:"sem_url_action"`
		} `json:"disable_autofilter"`
		ResetAutofilterfields struct {
			SemURL       string `json:"sem_url"`
			SemURLAction string `json:"sem_url_action"`
		} `json:"reset_autofilterfields"`
	} `json:"base_urls"`
	ResponseHeader struct {
		SemQTime  int `json:"semQTime"`
		QTime     int `json:"QTime"`
		SemStatus int `json:"semStatus"`
		Status    int `json:"status"`
		Params    struct {
			Q         string    `json:"q"`
			SphFq     time.Time `json:"sph.fq"`
			Sort      string    `json:"sort"`
			Rows      string    `json:"rows"`
			SphFl     string    `json:"sph.fl"`
			SphSfield string    `json:"sph.sfield"`
			JSONNl    string    `json:"json.nl"`
		} `json:"params"`
	} `json:"responseHeader"`
}
