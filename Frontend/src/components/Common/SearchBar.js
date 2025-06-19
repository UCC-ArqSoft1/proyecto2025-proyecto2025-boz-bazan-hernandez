// SearchBar.js - Versión actualizada con estilos modernos
import React, { useState, useEffect } from 'react';
import './SearchBar.css'; // Importar los estilos

const SearchBar = ({ onSearch }) => {
  const [query, setQuery] = useState('');
  const [suggestions, setSuggestions] = useState([]);
  const [showSuggestions, setShowSuggestions] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  // Sugerencias populares
  const popularTags = [
    { text: 'Fitness', icon: '💪' },
    { text: 'Yoga', icon: '🧘‍♀️' },
    { text: 'Natación', icon: '🏊‍♂️' },
    { text: 'Baile', icon: '💃' },
    { text: 'Deportes', icon: '⚽' },
    { text: 'Pilates', icon: '🤸‍♀️' },
    { text: 'Running', icon: '🏃‍♂️' },
    { text: 'Crossfit', icon: '🏋️‍♂️' }
  ];

  // Sugerencias de búsqueda simuladas
  const searchSuggestions = [
    { text: 'Yoga matutino', category: 'Yoga', icon: '🌅' },
    { text: 'Fitness nocturno', category: 'Fitness', icon: '🌙' },
    { text: 'Natación principiantes', category: 'Natación', icon: '🏊‍♀️' },
    { text: 'Baile latino', category: 'Baile', icon: '💃' },
    { text: 'Futbol 5', category: 'Deportes', icon: '⚽' },
    { text: 'Pilates avanzado', category: 'Pilates', icon: '🤸‍♀️' }
  ];

  const handleSubmit = (e) => {
    e.preventDefault();
    setIsLoading(true);
    onSearch(query);
    setShowSuggestions(false);
    
    // Simular loading
    setTimeout(() => {
      setIsLoading(false);
    }, 1000);
  };

  const handleInputChange = (e) => {
    const value = e.target.value;
    setQuery(value);
    
    if (value.length > 0) {
      const filtered = searchSuggestions.filter(
        suggestion => 
          suggestion.text.toLowerCase().includes(value.toLowerCase()) ||
          suggestion.category.toLowerCase().includes(value.toLowerCase())
      );
      setSuggestions(filtered);
      setShowSuggestions(true);
    } else {
      setShowSuggestions(false);
    }
  };

  const handleSuggestionClick = (suggestion) => {
    setQuery(suggestion.text);
    setShowSuggestions(false);
    onSearch(suggestion.text);
  };

  const handleTagClick = (tag) => {
    setQuery(tag.text);
    onSearch(tag.text);
  };

  const handleInputFocus = () => {
    if (query.length > 0) {
      setShowSuggestions(true);
    }
  };

  const handleInputBlur = () => {
    // Delay para permitir clicks en sugerencias
    setTimeout(() => {
      setShowSuggestions(false);
    }, 200);
  };

  return (
    <div className="search-container">
      <form onSubmit={handleSubmit} className="search-form">
        <div className="search-wrapper">
          <div className="search-icon">
            🔍
          </div>
          <input
            type="text"
            value={query}
            onChange={handleInputChange}
            onFocus={handleInputFocus}
            onBlur={handleInputBlur}
            placeholder="Buscar actividades por título, categoría o día..."
            className="search-input"
            disabled={isLoading}
          />
          <button
            type="submit"
            className="search-button"
            disabled={isLoading}
          >
            {isLoading ? (
              <>
                <div className="search-loading-spinner"></div>
                Buscando...
              </>
            ) : (
              <>
                <span className="search-button-icon">🚀</span>
                Buscar
              </>
            )}
          </button>
        </div>

        {/* Sugerencias dropdown */}
        {showSuggestions && suggestions.length > 0 && (
          <div className="search-suggestions">
            {suggestions.map((suggestion, index) => (
              <div
                key={index}
                className="suggestion-item"
                onClick={() => handleSuggestionClick(suggestion)}
              >
                <span className="suggestion-icon">{suggestion.icon}</span>
                <span className="suggestion-text">{suggestion.text}</span>
                <span className="suggestion-category">{suggestion.category}</span>
              </div>
            ))}
          </div>
        )}
      </form>

      {/* Tags populares */}
      <div className="search-tags">
        {popularTags.map((tag, index) => (
          <button
            key={index}
            className={`search-tag ${query === tag.text ? 'active' : ''}`}
            onClick={() => handleTagClick(tag)}
          >
            {tag.icon} {tag.text}
          </button>
        ))}
      </div>
    </div>
  );
};

export default SearchBar;