import React, { useState } from 'react';
import { SearchFilters } from './types';

interface SearchBarProps {
    onSearch: (filters: SearchFilters) => void;
    onClear: () => void;
}

const SearchBar: React.FC<SearchBarProps> = ({ onSearch, onClear }) => {
    const [filters, setFilters] = useState<SearchFilters>({
        keyword: '',
        categoria: '',
        horario: ''
    });

    const handleInputChange = (field: keyof SearchFilters, value: string) => {
        const newFilters = { ...filters, [field]: value };
        setFilters(newFilters);
        onSearch(newFilters);
    };

    const handleClear = () => {
        const emptyFilters = { keyword: '', categoria: '', horario: '' };
        setFilters(emptyFilters);
        onClear();
    };

    const categorias = [
        'Todas las categorías',
        'Cardio',
        'Funcional',
        'Fuerza',
        'Flexibilidad',
        'Deportes',
        'Boxeo',
        'MMA',
        'Spinning',
        'Yoga'
    ];

    const horarios = [
        'Todos los horarios',
        'Mañana (06:00-12:00)',
        'Tarde (12:00-18:00)',
        'Noche (18:00-22:00)'
    ];

    return (
        <div className="search-bar">
            <div className="search-bar-container">
                <div className="search-input-group">
                    <input
                        type="text"
                        placeholder="Buscar por título o profesor..."
                        value={filters.keyword || ''}
                        onChange={(e) => handleInputChange('keyword', e.target.value)}
                        className="search-input"
                    />
                    <i className="search-icon">🔍</i>
                </div>

                <div className="search-filters">
                    <select
                        value={filters.categoria || ''}
                        onChange={(e) => handleInputChange('categoria', e.target.value)}
                        className="search-select"
                    >
                        {categorias.map((categoria, index) => (
                            <option key={index} value={index === 0 ? '' : categoria}>
                                {categoria}
                            </option>
                        ))}
                    </select>

                    <select
                        value={filters.horario || ''}
                        onChange={(e) => handleInputChange('horario', e.target.value)}
                        className="search-select"
                    >
                        {horarios.map((horario, index) => (
                            <option key={index} value={index === 0 ? '' : horario}>
                                {horario}
                            </option>
                        ))}
                    </select>

                    <button onClick={handleClear} className="btn-clear" type="button">
                        Limpiar
                    </button>
                </div>
            </div>
        </div>
    );
};

export default SearchBar;
