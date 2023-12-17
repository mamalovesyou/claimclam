"use client";

import { useEffect, useState } from "react";
import type { ChangeEvent, FC } from "react";
import PropTypes from "prop-types";
import {
  Box,
  Checkbox,
  FormControl,
  FormControlLabel,
  InputAdornment,
  InputLabel,
  MenuItem,
  OutlinedInput,
  Select,
  TextField,
} from "@mui/material";
import { Search as SearchIcon } from "@mui/icons-material";
import { usePodcasts } from "@/hooks/usePodcasts";

type Filter = {
  label: string;
  value: string;
};

const filters: Filter[] = [
  {
    label: "All",
    value: "all",
  },
  {
    label: "Title",
    value: "title",
  },
  {
    label: "Category",
    value: "category",
  },
];

export const PodcastListFilters = () => {
  const [queryValue, setQueryValue] = useState<string>("");
  const [debouncedValue, setDebouncedValue] = useState<string>("");
  const [filter, setFilter] = useState<Filter>({
    label: "All",
    value: "all",
  });
  const [limit, setLimit] = useState<number>(10);

  const handleQueryChange = (event: ChangeEvent<HTMLInputElement>): void => {
    console.log("Query change", event.target.value);
    setQueryValue(event.target.value);
  };

  const { onFilterUpdate, pageInfo } = usePodcasts();

  useEffect(() => {
    const timeoutId = setTimeout(() => {
      setDebouncedValue(queryValue);
    }, 500);
    return () => clearTimeout(timeoutId);
  }, [queryValue]);

  useEffect(() => {
    console.log("useEffect", filter.value, limit, onFilterUpdate);
    if (!onFilterUpdate) return;
    console.log("onFilterUpdate", debouncedValue, filter.value);
    switch (filter.value) {
      case "title":
        onFilterUpdate({ title: debouncedValue, limit });
        console.log("onFilterUpdate title", { title: debouncedValue });
        break;
      case "category":
        onFilterUpdate({ category: debouncedValue, limit });
        console.log("onFilterUpdate category", { category: debouncedValue });
        break;
      default:
        onFilterUpdate({ search: debouncedValue, limit });
        console.log("onFilterUpdate search", { search: debouncedValue, limit });
        break;
    }
  }, [debouncedValue, filter.value, limit, onFilterUpdate]);

  return (
    <Box
      sx={{
        alignItems: "center",
        display: "flex",
      }}
    >
      <Box
        sx={{
          flexGrow: 1,
          ml: 3,
        }}
      >
        <TextField
          fullWidth
          variant="outlined"
          onChange={handleQueryChange}
          placeholder="Search a podcast..."
          value={queryValue}
          InputProps={{
            startAdornment: (
              <InputAdornment position="start">
                <SearchIcon />
              </InputAdornment>
            ),
          }}
        />
      </Box>
      <Box
        sx={{
          alignItems: "center",
          display: "flex",
          flexWrap: "wrap",
          p: 1,
        }}
      >
        <FormControl sx={{ m: 1, width: 300 }}>
          <InputLabel id="filters-select-label">Filters</InputLabel>
          <Select
            labelId="filters-select-label"
            id="filters-select"
            value={filter}
            input={<OutlinedInput label="Filters" />}
            renderValue={(filter) => filter.label}
            MenuProps={{
              PaperProps: {
                style: {
                  maxHeight: 185,
                  width: 175,
                },
              },
            }}
          >
            {filters.map((option) => (
              <MenuItem key={option.label}>
                <FormControlLabel
                  control={
                    <Checkbox
                      checked={filter.value === option.value}
                      onChange={() => setFilter(option)}
                      value={option.value}
                    />
                  }
                  label={option.label}
                  sx={{
                    flexGrow: 1,
                    mr: 0,
                  }}
                />
              </MenuItem>
            ))}
          </Select>
        </FormControl>
        <FormControl sx={{ m: 1, width: 125 }}>
          <InputLabel id="limit-select-label">Limits per page</InputLabel>
          <Select
            labelId="limit-select-label"
            id="limite-select"
            value={limit}
            label="Limits per page"
            onChange={(event) => setLimit(Number(event.target.value))}
          >
            <MenuItem value={10}>10</MenuItem>
            <MenuItem value={25}>25</MenuItem>
            <MenuItem value={50}>50</MenuItem>
          </Select>
        </FormControl>
      </Box>
    </Box>
  );
};

PodcastListFilters.propTypes = {
  onChange: PropTypes.func,
};
